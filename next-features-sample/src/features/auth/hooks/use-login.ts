"use client";

import { useState } from "react";
import { login } from "../api/login";
import type { User } from "../types";

/**
 * 画面の見た目を持たない、auth 機能の状態とふるまいだけを担当する hook。
 *
 * ここに手続きを寄せておくと、component 側は「表示」だけになる。
 * この hook は auth の内部実装なので、他 feature から import されない
 * (biome.json の境界ルールで機械的に禁止している)。
 */
export function useLogin() {
  const [user, setUser] = useState<User | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isPending, setIsPending] = useState(false);

  async function submit(email: string, password: string) {
    setIsPending(true);
    setError(null);
    try {
      setUser(await login({ email, password }));
    } catch (e) {
      setError(e instanceof Error ? e.message : "不明なエラーが発生しました");
    } finally {
      setIsPending(false);
    }
  }

  function logout() {
    setUser(null);
    setError(null);
  }

  return { user, error, isPending, submit, logout };
}
