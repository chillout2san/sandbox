"use client";

import { useState } from "react";
import { Button } from "@/components";
import { useTodos } from "../hooks/use-todos";
import styles from "./todo-panel.module.css";

export function TodoPanel() {
  const { todos, remaining, add, toggle, remove } = useTodos();
  const [draft, setDraft] = useState("");

  return (
    <section className={styles.panel}>
      <header className={styles.header}>
        <h2 className={styles.title}>Todo</h2>
        <span className={styles.remaining}>残り {remaining} 件</span>
      </header>

      <form
        className={styles.form}
        onSubmit={(e) => {
          e.preventDefault();
          add(draft);
          setDraft("");
        }}
      >
        <input
          className={styles.input}
          value={draft}
          onChange={(e) => setDraft(e.target.value)}
          placeholder="やることを入力"
          aria-label="やることを入力"
        />
        {/* auth と同じ共有 Button を、todo の文脈で使い回している */}
        <Button type="submit" disabled={!draft.trim()}>
          追加
        </Button>
      </form>

      {todos.length === 0 ? (
        <p className={styles.empty}>まだ Todo がありません</p>
      ) : (
        <ul className={styles.list}>
          {todos.map((todo) => (
            <li
              key={todo.id}
              className={`${styles.item} ${todo.done ? styles.done : ""}`}
            >
              <label className={styles.itemLabel}>
                <input
                  type="checkbox"
                  checked={todo.done}
                  onChange={() => toggle(todo.id)}
                />
                <span className={styles.itemText}>{todo.title}</span>
              </label>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => remove(todo.id)}
                aria-label={`${todo.title} を削除`}
              >
                削除
              </Button>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
