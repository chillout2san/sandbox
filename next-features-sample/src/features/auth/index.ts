/**
 * auth feature の公開 API。
 *
 * 外から使ってよいものだけをここに並べる。
 * use-login や api/login は意図的に export していない = 外から触れない内部実装。
 * この「窓口を 1 つに絞る」のが features 構成の実質的な効能で、
 * Lint 側では `@/features/auth/**` への深い import を禁止して裏付けている。
 */
export { LoginForm } from "./components/login-form";
export type { User } from "./types";
