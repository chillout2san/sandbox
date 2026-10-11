# next-features-sample

`pages` / `components` / `features` の 3 層構成を Next.js App Router で試すサンプル。
「構成を守る」ところを口約束にせず、**Biome の Lint で落ちるようにする**のが主眼。

```bash
npm install
npm run dev        # http://localhost:3000
npm run lint       # Biome: format + lint + 境界ルール
npm run format     # 自動修正
npm run typecheck  # tsc --noEmit
```

## ディレクトリ

```
src/
  app/                      ルーティング層 (質問の pages に相当)
    layout.tsx
    page.tsx
    login/page.tsx          → features/auth を組み立てるだけ
    todos/page.tsx          → features/todo を組み立てるだけ
    globals.css             デザイントークン + 最小リセット

  components/               機能に依存しない共有部品 (最下層)
    button/
      button.tsx
      button.module.css
      index.ts
    index.ts                公開 API

  features/
    auth/
      api/login.ts          通信 (バックエンドのリポジトリ層に近い)
      hooks/use-login.ts    状態とふるまい
      utils/format.ts       この feature だけで使うヘルパー
      components/login-form.tsx
      types.ts
      index.ts              公開 API ← 外から見えるのはここだけ
    todo/
      hooks/use-todos.ts
      components/todo-panel.tsx
      types.ts
      index.ts

scripts/
  verify-boundaries.sh      境界ルールが効いているかの検証 (npm run verify:boundaries)
```

`features/auth/` の中の `api` / `hooks` / `utils` / `components` という分け方は**規約ではない**。
この feature に必要なものを置いているだけで、`todo/` には `api/` も `utils/` もない。
境界ルールもサブディレクトリ名に依存していないので、feature ごとに好きに決めてよい。

依存の向きは常に **`app` → `features` → `components`** の一方通行。逆流と横移動（feature 間）を禁止する。

| レイヤー | 依存してよい | 禁止 |
| --- | --- | --- |
| `app/` | `features/*` の **index のみ**、`components` | feature の内部実装への直接 import |
| `features/<x>/` | 自分の中、`components` | 他の feature、`app` |
| `components/` | 自分の中のみ | すべての feature、`app` |

## 境界を Lint で守る

ESLint の `import/no-restricted-paths` に相当するものを **Biome の `style/noRestrictedImports` + `overrides`** で表現している（`biome.jsonc`）。
`overrides.includes` で「どのディレクトリのファイルか」を指定し、その中で禁止する import パターンを宣言する。

### 前提となる Biome の 3 つの性質

設定の形はこの 3 つから決まっている（すべて実測で確認済み）。

1. `noRestrictedImports` は **解決後のパスではなく import 文の文字列**にマッチする。
   そのためエイリアス形 (`@/features/todo/types`) と相対形 (`../todo/types`) の
   **両方**を塞ぐ必要がある。片方だけだともう片方が素通りする。
2. **overrides は同じルールに対してマージされず「後勝ちで全置換」される。**
   よって 1 ファイルに効かせたいパターンは 1 つの override に全部書く。分割できない。
3. **`includes` の `*` は `/` を跨がない。** `src/features/*/*` は feature 直下の
   ファイルだけにマッチする。つまり**ディレクトリの深さで場合分けできる**。

### 深さで場合分けするという解き方

素朴にやると「`../hooks/**` は許可、`../api/**` は許可、…」と**自分の中のサブディレクトリを
名前で列挙する**ことになり、`utils/` を 1 つ足した瞬間に正当な import が誤検知される。

そこで 3 を使い、**深さごとに「何階層上がったら feature の外に出るか」**を考える。
その 1 つ上を丸ごと禁止すればよく、サブディレクトリ名は設定に出てこない。

| ファイルの位置 | `../` の行き先 | 禁止するもの |
| --- | --- | --- |
| `features/auth/index.ts` | `features/` = **外** | `../**` |
| `features/auth/hooks/use-x.ts` | `auth/` = 中 | `../../**` |
| `features/auth/a/b/c.ts` | `auth/a/` = 中 | `../../../**` |

```jsonc
{
  // feature 内 1 階層目。`../` は自分の feature 内なので許可、`../../` 以上を禁止
  "includes": ["src/features/*/*/*"],
  "linter": {
    "rules": {
      "style": {
        "noRestrictedImports": {
          "level": "error",
          "options": {
            "patterns": [{
              "group": [
                "@/features/**",      // 他 feature も自分自身もエイリアス禁止 (中は相対で書く)
                "../../**",           // feature の外に出る相対参照を禁止
                "@/app/**",           // ルーティング層への逆流を禁止
                "@/components/*/**"   // 共有部品は index 経由のみ
              ],
              "message": "feature の外には出られません。"
            }]
          }
        }
      }
    }
  }
}
```

`@/features/**` を**自分の feature も含めて**全面禁止しているのがコツ。
「feature の中は相対パス、外に出るときだけエイリアス」という規約にすると、
feature 名が設定から消えるので**feature を増やしても設定変更が不要**になる。

結果として、**feature を増やしても、feature 内にサブディレクトリを増やしても、
`biome.jsonc` は一切変更しなくてよい**。

### 検証

その「増えても壊れない」ことを含めて `npm run verify:boundaries` で検証できる
（`scripts/verify-boundaries.sh`）。違反 13 件 + 正当な import 7 件を一時的に生成して
lint にかけ、期待どおり検出 / 非検出になるか確認して片付ける。

```
=== NG: 検出されるべき違反 ===
  ✅ ng01   NG  detected   features/auth/ng01.ts      → "../todo/types"
  ✅ ng02   NG  detected   features/auth/deep/        → "../../todo/types"
  ...
  ✅ ng12   NG  detected   features/payment/          → "../auth/types"     ★新 feature
=== OK: 検出されてはいけない正当な import ===
  ✅ ok01   OK  clean      features/auth/components/  → "../stores/store"   ★新サブディレクトリ
  ✅ ok07   OK  clean      features/auth/deep/nested/ → "../../utils/format"
pass=20 fail=0
```

★ の 2 件が肝で、**設定を触らずに**新しい feature が制約され、
新しいサブディレクトリが誤検知されないことを示している。

### 既知の限界

パス解決をしないので、**文字列として想定外の形は理屈上すり抜けうる**
（例: `require()` や動的 `import()`、tsconfig に別のエイリアスを足した場合）。
実際にモジュールを解決する `dependency-cruiser` / `eslint-plugin-boundaries` /
ESLint の `import/no-restricted-paths` のほうがこの点は厳密。
より強い保証が欲しくなったら `dependency-cruiser` を CI に足すのが素直（Biome と併用できる）。

なお **oxlint (v1.87 時点) には `import/no-restricted-paths` が未実装**で、
設定を書くと起動時に `Rule 'no-restricted-paths' not found in plugin 'import'` で落ちる。
`no-restricted-imports` は実装済みなので同じ書き方はできるが、
フォーマッタも兼ねる Biome のほうがこの用途では扱いやすい。

## UI ライブラリを使っていない理由

このサンプルは「大きめの会社が自前のデザインシステムを持つ」ノリに寄せてある。
`globals.css` のトークンを 2 層構成にしているのがその中核。

```
プリミティブ   --color-blue-600   生の値。意味を持たない
      ↓
セマンティック --accent-default   用途で名付ける。コンポーネントが参照するのはこちら
      ↓
コンポーネント button.module.css  var(--accent-default) しか書かない
```

### ヘッドレス UI について

今は「ゼロから全部書く」より **ヘッドレス UI（Radix UI / React Aria など）を土台に
自前のデザインシステムを載せる**のが主流。ヘッドレス UI は
**ふるまいとアクセシビリティだけを提供して、見た目を一切持たないライブラリ**で、
モーダル・セレクト・タブのような「自前で正しく作るのが大変なもの」だけを借りる。

借りられるのは、たとえばモーダルなら以下のような部分。

- `Esc` で閉じる、背景クリックで閉じる
- 開いたらフォーカスを中に移し、閉じたら元の要素に戻す
- Tab がモーダルの外に出ない（フォーカストラップ）
- 背景のスクロールを止める
- `role="dialog"` / `aria-modal` / ラベルの関連付け
- スクリーンリーダーに「ダイアログが開いた」と伝える

このサンプルは Button だけなのでヘッドレス UI を入れていない。
Button は `<button>` 要素がほぼすべてをやってくれるので、借りる必要がない部類。
逆に Modal や Select を作る段になったら入れる価値が出る。

### トークン層

`button.module.css` には生の色が 1 つも出てこない。この規律が守られている限り、
ブランド配色の変更やダークモード対応は**セマンティック層の差し替えだけ**で済む
（実際、ダークモードは `prefers-color-scheme` でセマンティック層を上書きしているだけ）。
