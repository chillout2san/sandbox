# CLAUDE.md

システムデザインの練習をするディレクトリ。お題ごとに設計し、最終成果物として **インフラ構成図** と **Terraform のコード** を作る。

## 構成

お題ごとにディレクトリを掘る（ハイフン区切り。例: `scale-up`）。

```
system-design/
├── .claude/skills/stack-diagram/  # 構成図用のスキル（npx skills add caru-ini/stack-diagram で導入）
├── skills-lock.json               # スキルのロックファイル
└── scale-up/
    ├── diagram.html               # インフラ構成図（1 つだけ）
    └── terraform/                 # Terraform のコード一式
```

## インフラ構成図

- `stack-diagram` スキル（`.claude/skills/stack-diagram/`）で作る
- スキルは `npx skills add caru-ini/stack-diagram -a claude-code -y` で導入した。更新は `npx skills update` で行い、手で編集しない
- お題ディレクトリの root に、ビルド済みの self-contained な HTML を 1 つだけ置く
- アイコンの取得先やビルド前の HTML などの中間ファイルはコミットしない（AWS 等の公式アイコンは再配布不可のため）

## Terraform

- お題ディレクトリ配下の `terraform/` にまとめる
- `.terraform/` や `*.tfstate` はコミットしない
