"use client";

import { Phrase } from "@/lib/types";

interface PhraseCardProps {
  phrase: Phrase;
  onPlay: (text: string) => void;
  isPlaying: boolean;
}

export function PhraseCard({ phrase, onPlay, isPlaying }: PhraseCardProps) {
  return (
    <div className="bg-white rounded-xl shadow-sm border border-gray-100 p-4 flex items-start gap-4">
      <div className="flex-1 min-w-0">
        <p className="text-sm text-gray-500 mb-1 truncate">{phrase.japaneseText}</p>
        <p className="text-base font-medium text-gray-900 leading-snug">{phrase.englishText}</p>
      </div>
      <button
        onClick={() => onPlay(phrase.englishText)}
        disabled={isPlaying}
        aria-label="Play audio"
        className="flex-shrink-0 w-10 h-10 rounded-full flex items-center justify-center transition-colors
          bg-emerald-100 text-emerald-700 hover:bg-emerald-200 disabled:opacity-50 disabled:cursor-not-allowed"
      >
        {isPlaying ? (
          <svg className="w-5 h-5 animate-pulse" fill="currentColor" viewBox="0 0 20 20">
            <rect x="5" y="5" width="3" height="10" rx="1" />
            <rect x="12" y="5" width="3" height="10" rx="1" />
          </svg>
        ) : (
          <svg className="w-5 h-5 ml-0.5" fill="currentColor" viewBox="0 0 20 20">
            <path d="M6.3 2.84A1.5 1.5 0 004 4.11v11.78a1.5 1.5 0 002.3 1.27l9.344-5.891a1.5 1.5 0 000-2.538L6.3 2.84z" />
          </svg>
        )}
      </button>
    </div>
  );
}
