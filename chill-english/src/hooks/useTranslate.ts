"use client";

import { useState } from "react";

interface UseTranslateResult {
  candidates: string[];
  isLoading: boolean;
  error: string | null;
  translate: (japaneseText: string) => Promise<void>;
  clearCandidates: () => void;
}

export function useTranslate(): UseTranslateResult {
  const [candidates, setCandidates] = useState<string[]>([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const translate = async (japaneseText: string) => {
    setIsLoading(true);
    setError(null);
    setCandidates([]);

    try {
      const res = await fetch("/api/translate", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ japaneseText }),
      });

      if (!res.ok) {
        throw new Error("Translation request failed");
      }

      const data = await res.json();
      setCandidates(data.candidates);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Translation failed");
    } finally {
      setIsLoading(false);
    }
  };

  const clearCandidates = () => {
    setCandidates([]);
    setError(null);
  };

  return { candidates, isLoading, error, translate, clearCandidates };
}
