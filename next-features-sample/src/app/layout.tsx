import type { Metadata } from "next";
import { Geist } from "next/font/google";
import Link from "next/link";
import "./globals.css";
import styles from "./layout.module.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "features ディレクトリ構成サンプル",
  description:
    "pages / components / features の 3 層構成を Next.js App Router で試すサンプル",
};

export default function RootLayout({ children }: LayoutProps<"/">) {
  return (
    <html lang="ja" className={geistSans.variable}>
      <body>
        <div className={styles.shell}>
          <nav className={styles.nav}>
            <Link href="/" className={styles.brand}>
              features sample
            </Link>
            <Link href="/login">ログイン</Link>
            <Link href="/todos">Todo</Link>
          </nav>
          <main>{children}</main>
        </div>
      </body>
    </html>
  );
}
