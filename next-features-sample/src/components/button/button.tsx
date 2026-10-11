import type { ButtonHTMLAttributes } from "react";
import styles from "./button.module.css";

/**
 * 共有レイヤーのボタン。
 *
 * ここは「どの機能からも使われる」側なので、auth や todo の事情を一切知らない。
 * ドメインの言葉 (「ログインボタン」「削除する Todo」) を型やデフォルト値に
 * 持ち込まないのが共有コンポーネントの境界線。
 */

type Variant = "primary" | "secondary" | "ghost" | "danger";
type Size = "sm" | "md" | "lg";

export type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  variant?: Variant;
  size?: Size;
  fullWidth?: boolean;
  /** true の間は disabled 扱いになり、スピナーを表示する */
  loading?: boolean;
};

export function Button({
  variant = "primary",
  size = "md",
  fullWidth = false,
  loading = false,
  disabled,
  className,
  children,
  ...rest
}: ButtonProps) {
  const classNames = [
    styles.base,
    styles[variant],
    styles[size],
    fullWidth ? styles.fullWidth : null,
    className,
  ]
    .filter(Boolean)
    .join(" ");

  return (
    <button
      type="button"
      className={classNames}
      disabled={disabled || loading}
      aria-busy={loading || undefined}
      {...rest}
    >
      {loading ? <span className={styles.spinner} aria-hidden="true" /> : null}
      {children}
    </button>
  );
}
