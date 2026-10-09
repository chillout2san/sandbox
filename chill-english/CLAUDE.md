# Chill English

シャドーイング練習用 Web アプリ。日本語を入力すると英訳候補を生成し、選択したフレーズを TTS で読み上げる。

## 機能

1. **日本語入力 → 英訳生成**: 入力した日本語を Gemini API に送り、自然な口語英文を 1 候補生成して表示する
2. **翻訳候補の試聴**: 候補カードの再生ボタンで TTS を再生し、選択前に発音を確認できる
3. **フレーズ保存**: 候補を選択すると Firestore のセッションに保存され、フレーズ一覧に追加される
4. **TTS 読み上げ**: フレーズ一覧の各カードから Google Cloud TTS で英文を読み上げる
5. **セッション永続化**: `sessionStorage` でセッション ID を管理し、リロード後もフレーズ一覧を復元する

## 技術スタック

| カテゴリ | 内容 |
|----------|------|
| フレームワーク | Next.js 16 (App Router) |
| 言語 | TypeScript |
| スタイリング | Tailwind CSS v4 |
| 翻訳 AI | Gemini (`gemini-2.5-flash`) via `@google/generative-ai` |
| TTS | Google Cloud TTS REST API (`en-US-Chirp3-HD-Kore`) |
| DB | Firebase Firestore (`firebase` SDK v12) |
| ランタイム | Node.js (Next.js サーバーサイド API Routes) |

## ディレクトリ構成

```
src/
├── app/
│   ├── page.tsx                  # メインページ・音声再生ロジック
│   ├── layout.tsx
│   └── api/
│       ├── translate/route.ts    # Gemini 翻訳エンドポイント
│       └── tts/route.ts          # Google Cloud TTS エンドポイント
├── components/
│   ├── JapaneseInputForm.tsx     # 日本語入力フォーム
│   ├── TranslationCandidates.tsx # 翻訳候補一覧（再生ボタン付き）
│   ├── PhraseList.tsx            # 保存フレーズ一覧
│   └── PhraseCard.tsx            # フレーズカード（再生ボタン付き）
├── hooks/
│   ├── useTranslate.ts           # 翻訳 API 呼び出し・状態管理
│   └── useSession.ts             # セッション初期化・フレーズ追加
└── lib/
    ├── firebase.ts               # Firebase 初期化
    ├── firestore.ts              # createSession / addPhrase / getPhrases
    └── types.ts                  # Phrase / Session 型定義
```

## Firestore データモデル

```
sessions/{sessionId}
  createdAt: Timestamp
  updatedAt: Timestamp
  phrases/{phraseId}
    japaneseText: string
    englishText: string
    order: number
    createdAt: Timestamp
```

## API エンドポイント

### POST /api/translate
- **Request**: `{ japaneseText: string }`
- **Response**: `{ candidates: string[] }`
- Gemini に口語英訳を 1 件生成させ JSON 配列で返す

### POST /api/tts
- **Request**: `{ text: string, speakingRate?: number }`
- **Response**: `{ audioContent: string }` (Base64 MP3)
- `speakingRate` のデフォルトは `1.0`

## 環境変数

```env
# サーバーサイドのみ
GEMINI_API_KEY=
GOOGLE_TTS_API_KEY=

# クライアント（Firebase）
NEXT_PUBLIC_FIREBASE_API_KEY=
NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN=
NEXT_PUBLIC_FIREBASE_PROJECT_ID=
NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET=
NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID=
NEXT_PUBLIC_FIREBASE_APP_ID=
```

## 課金

使用している外部サービスの課金情報。

| サービス | 無料枠 | 超過時の料金 | 備考 |
|----------|--------|-------------|------|
| **Gemini API** (`gemini-2.5-flash`) | 1日あたり一定リクエスト数まで無料（Free Tier） | 入力 $0.15/1M tokens、出力 $0.60/1M tokens（目安） | [料金詳細](https://ai.google.dev/pricing) |
| **Google Cloud TTS** | 月100万文字まで無料（WaveNet/Neural2系は40万文字） | $16/1M文字（HD Chirp系） | [料金詳細](https://cloud.google.com/text-to-speech/pricing) |
| **Firebase Firestore** | 読み取り5万回/書き込み2万回/削除2万回（1日あたり） | $0.06/10万読み取り など | [料金詳細](https://firebase.google.com/pricing) |

> 個人利用・開発用途であれば基本的に無料枠内に収まる想定。

## 開発サーバー起動

```bash
npm run dev
```
