"use client";

import { useState } from "react";
import type { Todo } from "../types";

const initialTodos: Todo[] = [
  { id: "t_1", title: "features の切り方を理解する", done: true },
  { id: "t_2", title: "Biome で境界ルールを書く", done: false },
  { id: "t_3", title: "Button をデザイントークンで組む", done: false },
];

export function useTodos() {
  const [todos, setTodos] = useState<Todo[]>(initialTodos);

  function add(title: string) {
    const trimmed = title.trim();
    if (!trimmed) return;
    setTodos((prev) => [
      ...prev,
      { id: `t_${crypto.randomUUID()}`, title: trimmed, done: false },
    ]);
  }

  function toggle(id: string) {
    setTodos((prev) =>
      prev.map((todo) => (todo.id === id ? { ...todo, done: !todo.done } : todo)),
    );
  }

  function remove(id: string) {
    setTodos((prev) => prev.filter((todo) => todo.id !== id));
  }

  const remaining = todos.filter((todo) => !todo.done).length;

  return { todos, remaining, add, toggle, remove };
}
