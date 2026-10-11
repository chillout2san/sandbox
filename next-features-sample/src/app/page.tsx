import Link from "next/link";
import styles from "./page.module.css";

const tree = `src/
  app/          … ルーティング層。features を組み立てるだけ
  components/   … 機能に依存しない共有部品 (Button)
  features/
    auth/       … components / hooks / api / types + index.ts
    todo/       … components / hooks / types + index.ts`;

export default function Home() {
  return (
    <div className={styles.hero}>
      <h1 className={styles.title}>pages / components / features の 3 層構成</h1>
      <p className={styles.lead}>
        役割の違う 3 種類を分けているだけの構成です。feature 同士が直接依存しないことを
        Biome の <code>noRestrictedImports</code> で機械的に保証しています。
      </p>
      <pre className={styles.tree}>{tree}</pre>
      <div className={styles.links}>
        <Link href="/login">ログイン画面へ</Link>
        <Link href="/todos">Todo 画面へ</Link>
      </div>
    </div>
  );
}
