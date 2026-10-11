#!/bin/bash
# 境界ルールの網羅検証。
# NG = 検出されるべき / OK = 検出されてはいけない
cd /Users/hayashi/ghq/github.com/chillout2san/sandbox/next-features-sample || exit 1

TMP=".verify"
rm -rf "$TMP"; mkdir -p "$TMP"

# --- 検証用の一時ファイル群を作る ---
mkdir -p src/features/auth/deep/nested
mkdir -p src/features/payment/components     # ★ 新しい feature (設定は未変更)
mkdir -p src/features/auth/stores            # ★ 新しいサブディレクトリ
mkdir -p src/components/card                 # ★ 新しい共有部品

printf 'export const s = 1;\n'                                             > src/features/auth/stores/store.ts
printf 'export const Card = 1;\n'                                          > src/components/card/card.ts
printf 'export type Pay = { id: string };\n'                               > src/features/payment/types.ts

# NG ケース
printf 'import type { Todo } from "../todo/types";\nexport const x: Todo | null = null;\n'                      > src/features/auth/ng01.ts
printf 'import type { Todo } from "../../todo/types";\nexport const x: Todo | null = null;\n'                   > src/features/auth/deep/ng02.ts
printf 'import type { Todo } from "../../../todo/types";\nexport const x: Todo | null = null;\n'                > src/features/auth/deep/nested/ng03.ts
printf 'import type { Todo } from "@/features/todo";\nexport const x: Todo | null = null;\n'                    > src/features/auth/ng04.ts
printf 'import type { Todo } from "@/features/todo/types";\nexport const x: Todo | null = null;\n'              > src/features/auth/ng05.ts
printf 'import Page from "@/app/page";\nexport const x = Page;\n'                                               > src/features/auth/ng06.ts
printf 'import { Button } from "@/components/button/button";\nexport const x = Button;\n'                       > src/features/auth/ng07.ts
printf 'import type { User } from "@/features/auth";\nexport const x: User | null = null;\n'                    > src/components/ng08.ts
printf 'import { LoginForm } from "@/features/auth/components/login-form";\nexport const x = LoginForm;\n'      > src/app/ng09.ts
printf 'import { Button } from "@/components/button/button";\nexport const x = Button;\n'                       > src/app/ng10.ts
printf 'import { LoginForm } from "../../features/auth/components/login-form";\nexport const x = LoginForm;\n'  > src/app/login/ng11.ts
printf 'import type { User } from "../auth/types";\nexport const x: User | null = null;\n'                      > src/features/payment/ng12.ts
printf 'import type { User } from "../../auth/types";\nexport const x: User | null = null;\n'                   > src/features/payment/components/ng13.ts

# OK ケース
printf 'import { s } from "../stores/store";\nexport const x = s;\n'                       > src/features/auth/components/ok01.ts
printf 'import { Button } from "@/components";\nexport const x = Button;\n'                > src/features/payment/components/ok02.ts
printf 'import type { Pay } from "../types";\nexport const x: Pay | null = null;\n'        > src/features/payment/components/ok03.ts
printf 'import { LoginForm } from "@/features/auth";\nexport const x = LoginForm;\n'       > src/app/ok04.ts
printf 'import { Button } from "@/components";\nexport const x = Button;\n'                > src/app/ok05.ts
printf 'import { s } from "./stores/store";\nexport const x = s;\n'                        > src/features/auth/ok06.ts
printf 'import { maskEmail } from "../../utils/format";\nexport const x = maskEmail;\n'    > src/features/auth/deep/nested/ok07.ts

# --- lint 実行 ---
npx biome lint src > "$TMP/out.txt" 2>&1
grep -oE "^src/[a-zA-Z0-9/._-]+" "$TMP/out.txt" | grep -oE "(ng|ok)[0-9]+" | sort -u > "$TMP/flagged.txt"

pass=0; fail=0
check() { # $1=id  $2=NG|OK
  if grep -qx "$1" "$TMP/flagged.txt"; then got="detected"; else got="clean"; fi
  if [ "$2" = "NG" ]; then want="detected"; else want="clean"; fi
  if [ "$got" = "$want" ]; then printf "  ✅ %-6s %-3s %s\n" "$1" "$2" "$got"; pass=$((pass+1));
  else printf "  ❌ %-6s %-3s expected=%s got=%s\n" "$1" "$2" "$want" "$got"; fail=$((fail+1)); fi
}

echo "=== NG: 検出されるべき違反 ==="
for i in 01 02 03 04 05 06 07 08 09 10 11 12 13; do check "ng$i" NG; done
echo "=== OK: 検出されてはいけない正当な import ==="
for i in 01 02 03 04 05 06 07; do check "ok$i" OK; done
echo
echo "pass=$pass fail=$fail"

# --- 片付け ---
rm -f src/features/auth/ng0*.ts src/features/auth/ng1*.ts src/features/auth/ok06.ts
rm -rf src/features/auth/deep src/features/payment src/features/auth/stores src/components/card
rm -f src/components/ng08.ts src/app/ng09.ts src/app/ng10.ts src/app/login/ng11.ts
rm -f src/app/ok04.ts src/app/ok05.ts src/features/auth/components/ok01.ts
rm -rf "$TMP"
exit $fail
