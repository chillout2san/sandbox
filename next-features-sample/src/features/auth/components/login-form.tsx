"use client";

import { useState } from "react";
// 共有レイヤーは「上から下」への依存なので OK。公開 API (`@/components`) 経由で取る
import { Button } from "@/components";
import { useLogin } from "../hooks/use-login";
// `utils/` は後から足したディレクトリ。深さベースの境界ルールなので
// 新しいサブディレクトリを作っても biome.jsonc は変更不要。
import { maskEmail } from "../utils/format";
import styles from "./login-form.module.css";

export function LoginForm() {
  const { user, error, isPending, submit, logout } = useLogin();
  const [email, setEmail] = useState("sato@example.com");
  const [password, setPassword] = useState("password");

  if (user) {
    return (
      <div className={styles.card}>
        <div className={styles.signedIn}>
          <span className={styles.name}>{user.name} さんでログイン中</span>
          <span className={styles.email}>{maskEmail(user.email)}</span>
          {/* 共有 Button を auth の文脈で使う */}
          <Button variant="secondary" size="sm" onClick={logout}>
            ログアウト
          </Button>
        </div>
      </div>
    );
  }

  return (
    <form
      className={styles.card}
      onSubmit={(e) => {
        e.preventDefault();
        void submit(email, password);
      }}
    >
      <div className={styles.field}>
        <label className={styles.label} htmlFor="email">
          メールアドレス
        </label>
        <input
          id="email"
          className={styles.input}
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          autoComplete="email"
        />
      </div>

      <div className={styles.field}>
        <label className={styles.label} htmlFor="password">
          パスワード
        </label>
        <input
          id="password"
          className={styles.input}
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          autoComplete="current-password"
        />
      </div>

      {error ? <p className={styles.error}>{error}</p> : null}

      <div className={styles.actions}>
        <Button type="submit" loading={isPending} fullWidth>
          ログイン
        </Button>
      </div>
    </form>
  );
}
