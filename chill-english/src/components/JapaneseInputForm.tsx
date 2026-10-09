"use client";

import { useState, useRef, FormEvent } from "react";

interface JapaneseInputFormProps {
  onTranslate: (text: string) => void;
  isLoading: boolean;
}

export function JapaneseInputForm({
  onTranslate,
  isLoading,
}: JapaneseInputFormProps) {
  const [value, setValue] = useState("");
  const [isListening, setIsListening] = useState(false);
  const recognitionRef = useRef<SpeechRecognition | null>(null);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();
    const trimmed = value.trim();
    if (!trimmed) return;
    onTranslate(trimmed);
  };

  const toggleListening = () => {
    const SpeechRecognition =
      window.SpeechRecognition || window.webkitSpeechRecognition;
    if (!SpeechRecognition) {
      alert("このブラウザは音声入力に対応していません。");
      return;
    }

    if (isListening) {
      recognitionRef.current?.stop();
      return;
    }

    const recognition = new SpeechRecognition();
    recognition.lang = "ja-JP";
    recognition.interimResults = false;
    recognition.continuous = false;
    recognitionRef.current = recognition;

    recognition.onresult = (event) => {
      const transcript = event.results[0][0].transcript;
      setValue((prev) => prev + transcript);
    };
    recognition.onerror = () => setIsListening(false);
    recognition.onend = () => setIsListening(false);

    recognition.start();
    setIsListening(true);
  };

  return (
    <form onSubmit={handleSubmit} className="flex flex-col gap-2">
      <div className="relative">
        <textarea
          value={value}
          onChange={(e) => setValue(e.target.value)}
          placeholder="日本語を入力..."
          disabled={isLoading}
          rows={1}
          className="w-full px-4 py-3 pr-12 rounded-xl border border-gray-200 bg-white text-sm text-gray-900
            focus:outline-none focus:ring-2 focus:ring-emerald-400 focus:border-transparent
            placeholder-gray-400 disabled:opacity-60 resize-y"
        />
        <button
          type="button"
          onClick={toggleListening}
          disabled={isLoading}
          title={isListening ? "録音停止" : "音声入力"}
          className={`absolute right-2 top-2.5 p-1.5 rounded-lg transition-colors
            ${isListening
              ? "text-red-500 bg-red-50 hover:bg-red-100 animate-pulse"
              : "text-gray-400 hover:text-emerald-500 hover:bg-emerald-50"
            } disabled:opacity-40 disabled:cursor-not-allowed`}
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="currentColor" className="w-5 h-5">
            <path d="M8.25 4.5a3.75 3.75 0 1 1 7.5 0v8.25a3.75 3.75 0 1 1-7.5 0V4.5Z" />
            <path d="M6 10.5a.75.75 0 0 1 .75.75v1.5a5.25 5.25 0 1 0 10.5 0v-1.5a.75.75 0 0 1 1.5 0v1.5a6.751 6.751 0 0 1-6 6.709v2.291h3a.75.75 0 0 1 0 1.5h-7.5a.75.75 0 0 1 0-1.5h3v-2.291a6.751 6.751 0 0 1-6-6.709v-1.5A.75.75 0 0 1 6 10.5Z" />
          </svg>
        </button>
      </div>
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
