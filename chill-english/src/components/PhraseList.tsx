"use client";

import { Phrase } from "@/lib/types";
import { PhraseCard } from "./PhraseCard";

interface PhraseListProps {
  phrases: Phrase[];
  playingId: string | null;
  onPlay: (text: string, id: string) => void;
}

export function PhraseList({ phrases, playingId, onPlay }: PhraseListProps) {
  if (phrases.length === 0) {
    return (
      <div className="text-center py-12 text-gray-400">
        <p className="text-sm">まだフレーズがありません</p>
        <p className="text-xs mt-1">日本語を入力して翻訳してみましょう</p>
      </div>
    );
  }

  return (
    <div className="space-y-3">
      {phrases.map((phrase) => (
        <PhraseCard
          key={phrase.id}
          phrase={phrase}
          isPlaying={playingId === phrase.id}
          onPlay={(text) => onPlay(text, phrase.id)}
        />
      ))}
    </div>
  );
}
