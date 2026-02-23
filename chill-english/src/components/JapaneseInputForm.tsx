"use client";

import { useState, FormEvent } from "react";

interface JapaneseInputFormProps {
  onTranslate: (text: string) => void;
  isLoading: boolean;
}

export function JapaneseInputForm({
  onTranslate,
  isLoading,
}: JapaneseInputFormProps) {
  const [value, setValue] = useState("");

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    const trimmed = value.trim();
    if (!trimmed) return;
    onTranslate(trimmed);
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-2">
      <textarea
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="日本語を入力..."
        disabled={isLoading}
        rows={1}
        className="w-full px-4 py-3 rounded-xl border border-gray-200 bg-white text-sm text-gray-900
          focus:outline-none focus:ring-2 focus:ring-emerald-400 focus:border-transparent
          placeholder-gray-400 disabled:opacity-60 resize-y"
      />
      <button
        type="submit"
        disabled={isLoading || !value.trim()}
        className="self-end px-5 py-2.5 rounded-xl bg-emerald-500 text-white text-sm font-medium
          hover:bg-emerald-600 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
      >
        {isLoading ? "翻訳中..." : "翻訳"}
      </button>
    </form>
  );
}
