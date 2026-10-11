import type { LoginInput, User } from "../types";

/**
 * 通信層。サンプルなので実際のエンドポイントは叩かず、擬似的に待つだけ。
 *
 * feature 内に api/ を置いておくと「この機能がサーバーと何をやり取りするか」が
 * 1 ディレクトリに閉じる。バックエンドで言えばリポジトリ層に近い役割。
 */
export async function login({ email, password }: LoginInput): Promise<User> {
  await new Promise((resolve) => setTimeout(resolve, 600));

  if (!email.includes("@") || password.length < 4) {
    throw new Error("メールアドレスまたはパスワードが正しくありません");
  }

  return {
    id: "u_1",
    name: email.split("@")[0] ?? "user",
    email,
  };
}
