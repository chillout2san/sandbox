BUF := go tool -modfile=tools/go.mod buf

# buf breaking の比較対象。既定は main ブランチの同ディレクトリ。
# make proto-breaking BREAKING_AGAINST=... で上書きできる。
# ('#' は make のコメント開始文字なのでエスケープが必要)
BREAKING_AGAINST ?= ../.git\#branch=main,subdir=connect-sample

.PHONY: proto-gen proto-lint proto-format proto-breaking proto-stats

# proto から Go のコードを生成する。
proto-gen:
	$(BUF) generate

# proto の lint をかける。
proto-lint:
	$(BUF) lint

# proto を整形して書き戻す。
proto-format:
	$(BUF) format -w

# 比較対象から見て後方互換性が壊れていないか検査する。
proto-breaking:
	$(BUF) breaking --against '$(BREAKING_AGAINST)'

# proto の統計を表示する。
proto-stats:
	$(BUF) stats
