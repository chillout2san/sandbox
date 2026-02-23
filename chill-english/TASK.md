# セキュリティ修正タスク（Vercel デプロイ前）

> このファイルは Vercel デプロイ前に対応すべきセキュリティ修正の一覧です。
> すべてのタスクが完了してから本番デプロイを行ってください。

---

## ステータス

- [ ] 🔴 **[Critical]** API ルートに認証を追加する
- [ ] 🔴 **[Critical]** Firestore セキュリティルールを設定する
- [ ] 🟠 **[High]** 入力テキストに長さ上限を追加する
- [ ] 🟠 **[High]** レートリミットを実装する
- [ ] 🟠 **[High]** Prompt Injection 対策を行う
- [ ] 🟠 **[High]** `JSON.parse` を try-catch で囲む
- [ ] 🟡 **[Medium]** CSP ヘッダを設定する
- [ ] 🟡 **[Medium]** `speakingRate` のバリデーションを追加する

---

## 詳細

### 🔴 [Critical] API ルートに認証を追加する

**対象ファイル:**
- `src/app/api/translate/route.ts`
- `src/app/api/tts/route.ts`

**問題:**
`/api/translate` と `/api/tts` は現在、ログインなしで誰でも呼び出せる。
UI は Firebase Auth で保護されているが、API エンドポイント自体は素通し。
Vercel にデプロイすると URL が公開され、第三者が Gemini / TTS を無制限に叩いて課金が発生するリスクがある。

**対応方針:**
1. Firebase Admin SDK をサーバーサイドに追加する（`firebase-admin` パッケージ）
2. クライアント側でログイン後に `user.getIdToken()` を取得し、リクエストヘッダに付与する
3. API ルート側でトークンを検証する

```typescript
// クライアント側（useTranslate.ts など）
const token = await user.getIdToken();
fetch("/api/translate", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    "Authorization": `Bearer ${token}`,
  },
  body: JSON.stringify({ japaneseText }),
});

// サーバー側（route.ts）
import { getApps, initializeApp, cert } from "firebase-admin/app";
import { getAuth } from "firebase-admin/auth";

const authHeader = req.headers.get("authorization");
if (!authHeader?.startsWith("Bearer ")) {
  return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
}
const token = authHeader.slice(7);
await getAuth().verifyIdToken(token); // 失敗すれば例外でエラー返却
```

---

### 🔴 [Critical] Firestore セキュリティルールを設定する

**対象:** Firebase Console → Firestore → ルール

**問題:**
現在のルールがテストモード（`allow read, write: if true`）の場合、
誰でも全データを読み書きできる。

**対応方針:**
Firebase Console でルールを以下に更新する。
ログイン済みの本人のセッション・フレーズのみ操作可能にする。

```
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /sessions/{sessionId} {
      allow read, write: if request.auth != null
                         && request.auth.uid == resource.data.userId;
      allow create:      if request.auth != null
                         && request.auth.uid == request.resource.data.userId;

      match /phrases/{phraseId} {
        allow read, write: if request.auth != null
                           && request.auth.uid ==
                              get(/databases/$(database)/documents/sessions/$(sessionId)).data.userId;
      }
    }
  }
}
```

---

### 🟠 [High] 入力テキストに長さ上限を追加する

**対象ファイル:**
- `src/app/api/translate/route.ts`
- `src/app/api/tts/route.ts`

**問題:**
現在テキスト長の上限がなく、巨大なテキストを送ると Gemini / TTS の費用が増加する。

**対応方針:**

```typescript
// translate/route.ts
const MAX_LENGTH = 500; // 日本語 500 文字が目安
if (japaneseText.length > MAX_LENGTH) {
  return NextResponse.json({ error: "Text too long" }, { status: 400 });
}

// tts/route.ts
const MAX_LENGTH = 1000; // 英文 1000 文字が目安
if (text.length > MAX_LENGTH) {
  return NextResponse.json({ error: "Text too long" }, { status: 400 });
}
```

---

### 🟠 [High] レートリミットを実装する

**対象ファイル:**
- `src/app/api/translate/route.ts`
- `src/app/api/tts/route.ts`

**問題:**
同一ユーザー / IP から短時間に大量リクエストを送られると、APIコストが増大する。

**対応方針:**
Upstash Redis + `@upstash/ratelimit` が Vercel と相性がよく手軽。

```bash
npm install @upstash/ratelimit @upstash/redis
```

```typescript
import { Ratelimit } from "@upstash/ratelimit";
import { Redis } from "@upstash/redis";

const ratelimit = new Ratelimit({
  redis: Redis.fromEnv(),
  limiter: Ratelimit.slidingWindow(20, "1 h"), // 1時間に 20 回まで
});

// route.ts の先頭で
const ip = req.headers.get("x-forwarded-for") ?? "unknown";
const { success } = await ratelimit.limit(ip);
if (!success) {
  return NextResponse.json({ error: "Rate limit exceeded" }, { status: 429 });
}
```

環境変数に `UPSTASH_REDIS_REST_URL` と `UPSTASH_REDIS_REST_TOKEN` を追加する。

---

### 🟠 [High] Prompt Injection 対策を行う

**対象ファイル:** `src/app/api/translate/route.ts`

**問題:**

```typescript
// 現在のコード（route.ts:25）
const prompt = `...Japanese text: "${japaneseText}"...`;
```

ユーザー入力がプロンプト文字列に直接埋め込まれており、
`"` や改行を含む入力でプロンプト構造を壊せる。

**対応方針:**
入力テキストをプロンプト本文から分離する。

```typescript
// 修正後：プロンプトとユーザー入力を分離
const systemPrompt = `
You are an English translation assistant for Japanese learners practicing shadowing.
Translate the given Japanese text into a single English sentence that sounds completely natural in everyday conversation.
Prioritize natural spoken rhythm, common contractions, and colloquial word choice over literal accuracy.
Return ONLY a JSON array containing exactly 1 string, no explanation, no markdown code blocks.
Example: ["translation"]
`;

const result = await model.generateContent([
  { text: systemPrompt },
  { text: `Japanese text to translate:\n${japaneseText}` },
]);
```

---

### 🟠 [High] `JSON.parse` を try-catch で囲む

**対象ファイル:** `src/app/api/translate/route.ts`

**問題:**

```typescript
// 現在のコード（route.ts:43）
const candidates: string[] = JSON.parse(match[0]); // ← try-catch なし
```

Gemini が不正な JSON を返したとき、例外がハンドルされず 500 エラーになる。

**対応方針:**

```typescript
let candidates: string[];
try {
  const parsed = JSON.parse(match[0]);
  if (!Array.isArray(parsed) || !parsed.every((v) => typeof v === "string")) {
    throw new Error("Unexpected format");
  }
  candidates = parsed;
} catch {
  return NextResponse.json(
    { error: "Failed to parse translation response" },
    { status: 500 }
  );
}
```

---

### 🟡 [Medium] CSP ヘッダを設定する

**対象ファイル:** `next.config.ts`

**問題:**
セキュリティヘッダが未設定で、XSS などへの多層防御がない。

**対応方針:**

```typescript
// next.config.ts
const nextConfig: NextConfig = {
  async headers() {
    return [
      {
        source: "/:path*",
        headers: [
          {
            key: "X-Frame-Options",
            value: "DENY",
          },
          {
            key: "X-Content-Type-Options",
            value: "nosniff",
          },
          {
            key: "Referrer-Policy",
            value: "strict-origin-when-cross-origin",
          },
          {
            key: "Content-Security-Policy",
            value: [
              "default-src 'self'",
              "script-src 'self' 'unsafe-inline' 'unsafe-eval'",
              "style-src 'self' 'unsafe-inline'",
              "connect-src 'self' https://*.googleapis.com https://*.google.com",
              "frame-src 'none'",
              "object-src 'none'",
            ].join("; "),
          },
        ],
      },
    ];
  },
};
```

---

### 🟡 [Medium] `speakingRate` のバリデーションを追加する

**対象ファイル:** `src/app/api/tts/route.ts`

**問題:**
`speakingRate` が Google TTS の許容範囲（0.25 〜 4.0）外の値でも
そのまま API に渡され、エラーになる。

**対応方針:**

```typescript
const speakingRate = typeof rawRate === "number"
  ? Math.min(Math.max(rawRate, 0.25), 4.0)
  : 1.0;
```

---

## 完了後のデプロイ手順

1. 上記タスクをすべて完了させる
2. `npm audit` を実行して既知の脆弱性がないか確認する
3. Vercel の環境変数に `.env.local` の値をすべて設定する
4. Firebase Console で Firestore セキュリティルールを適用する
5. GCP Console で各 API キーに適用サービスを制限する（任意・推奨）
6. `vercel deploy` を実行する
