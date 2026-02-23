"use client";

interface TranslationCandidatesProps {
  candidates: string[];
  onSelect: (candidate: string) => void;
  onPlay: (text: string, id: string) => void;
  playingId: string | null;
}

export function TranslationCandidates({
  candidates,
  onSelect,
  onPlay,
  playingId,
}: TranslationCandidatesProps) {
  if (candidates.length === 0) return null;

  return (
    <div className="mt-4">
      <p className="text-xs font-semibold text-gray-500 uppercase tracking-wide mb-2">
        翻訳候補を選択
      </p>
      <div className="space-y-2">
        {candidates.map((candidate, i) => {
          const id = `candidate-${i}`;
          const isPlaying = playingId === id;
          return (
            <div
              key={i}
              className="flex items-center gap-2 rounded-lg border border-gray-200 bg-white
                hover:border-emerald-400 hover:bg-emerald-50 transition-colors"
            >
              <button
                onClick={() => onSelect(candidate)}
                className="flex-1 text-left px-4 py-3 text-sm text-gray-800"
              >
                {candidate}
              </button>
              <button
                onClick={() => onPlay(candidate, id)}
                disabled={!!playingId}
                className="pr-3 text-gray-400 hover:text-emerald-500 disabled:opacity-30 transition-colors"
                aria-label="再生"
              >
                {isPlaying ? (
                  <svg className="w-5 h-5 animate-pulse" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z" />
                  </svg>
                ) : (
                  <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M8 5v14l11-7z" />
                  </svg>
                )}
              </button>
            </div>
          );
        })}
      </div>
    </div>
  );
}
