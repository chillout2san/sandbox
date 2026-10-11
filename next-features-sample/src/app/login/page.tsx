// feature は公開 API 経由でのみ import する。
// `@/features/auth/components/login-form` のような深い import は Biome が落とす。
import { LoginForm } from "@/features/auth";

export default function LoginPage() {
  return <LoginForm />;
}
