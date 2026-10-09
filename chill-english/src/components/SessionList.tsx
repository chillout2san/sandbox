"use client";

import { useState } from "react";
import { Session, Phrase } from "@/lib/types";
import { PhraseCard } from "./PhraseCard";

const PAGE_SIZE = 5;

interface SessionCardProps {
  session: Session;
  isActive: boolean;
  playingId: string | null;
  onSelectSession: (id: string) => void;
  onPlay: (text: string, id: string) => void;
  onMergePhrases: (sessionId: string) => Promise<void>;
}

function SessionCard({ session, isActive, playingId, onSelectSession, onPlay, onMergePhrases }: SessionCardProps) {
  const [phrasePage, setPhrasePage] = useState(0);
  const phrases: Phrase[] = session.phrases ?? [];
  const totalPhrasePages = Math.ceil(phrases.length / PAGE_SIZE);
  const pagedPhrases = phrases.slice(phrasePage * PAGE_SIZE, (phrasePage + 1) * PAGE_SIZE);
  const wordCount = phrases.reduce(
    (sum, p) => sum + p.englishText.trim().split(/\s+/).filter(Boolean).length,
    0
  );

  return (
    <div
      onClick={() => onSelectSession(session.id)}
      className={`rounded-2xl border p-4 cursor-pointer transition-colors ${
        isActive
          ? "border-emerald-400 bg-emerald-50"
          : "border-gray-200 bg-white hover:border-gray-300"
      }`}
    >
      <div className="flex items-center justify-between mb-3">
        <div className="flex items-center gap-2">
          <p className="text-sm font-semibold text-gray-800">{session.name}</p>
          <span className="text-xs text-gray-400">{wordCount} words</span>
          {isActive && (
            <span className="text-xs text-emerald-600 font-medium">選択中</span>
          )}
        </div>
        <div className="flex items-center gap-3" onClick={(e) => e.stopPropagation()}>
          {phrases.length >= 2 && (
            <button
              onClick={() => onMergePhrases(session.id)}
              className="text-xs text-gray-400 hover:text-gray-600 transition-colors"
            >
              結合
            </button>
          )}
          <span className="text-xs text-gray-400">{phrases.length} 件</span>
        </div>
      </div>

      {pagedPhrases.length > 0 && (
        <div className="space-y-2" onClick={(e) => e.stopPropagation()}>
          {pagedPhrases.map((phrase) => (
            <PhraseCard
              key={phrase.id}
              phrase={phrase}
              isPlaying={playingId === phrase.id}
              onPlay={(text) => onPlay(text, phrase.id)}
            />
          ))}

          {totalPhrasePages > 1 && (
            <div className="flex items-center justify-center gap-4 pt-1">
              <button
                onClick={() => setPhrasePage((p) => p - 1)}
                disabled={phrasePage === 0}
                className="text-xs text-gray-500 hover:text-gray-700 disabled:opacity-30 transition-colors"
              >
                ← 前へ
              </button>
              <span className="text-xs text-gray-400">
                {phrasePage + 1} / {totalPhrasePages}
              </span>
              <button
                onClick={() => setPhrasePage((p) => p + 1)}
                disabled={phrasePage === totalPhrasePages - 1}
                className="text-xs text-gray-500 hover:text-gray-700 disabled:opacity-30 transition-colors"
              >
                次へ →
              </button>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

interface SessionListProps {
  sessions: Session[];
  activeSessionId: string | null;
  playingId: string | null;
  onSelectSession: (id: string) => void;
  onPlay: (text: string, id: string) => void;
  onMergePhrases: (sessionId: string) => Promise<void>;
}

export function SessionList({
  sessions,
  activeSessionId,
  playingId,
  onSelectSession,
  onPlay,
  onMergePhrases,
}: SessionListProps) {
  const [page, setPage] = useState(0);
  const totalPages = Math.ceil(sessions.length / PAGE_SIZE);
  const paged = sessions.slice(page * PAGE_SIZE, (page + 1) * PAGE_SIZE);

  if (sessions.length === 0) {
    return (
      <div className="text-center py-12 text-gray-400">
        <p className="text-sm">セッションがありません</p>
        <p className="text-xs mt-1">新規セッションを作成してください</p>
      </div>
    );
  }

  return (
    <div>
      <div className="space-y-4">
        {paged.map((session) => (
          <SessionCard
            key={session.id}
            session={session}
            isActive={session.id === activeSessionId}
            playingId={playingId}
            onSelectSession={onSelectSession}
            onPlay={onPlay}
            onMergePhrases={onMergePhrases}
          />
        ))}
      </div>

      {totalPages > 1 && (
        <div className="flex items-center justify-center gap-4 mt-6">
          <button
            onClick={() => setPage((p) => p - 1)}
            disabled={page === 0}
            className="text-sm text-gray-500 hover:text-gray-700 disabled:opacity-30 transition-colors"
          >
            ← 前へ
          </button>
          <span className="text-xs text-gray-400">
            {page + 1} / {totalPages}
          </span>
          <button
            onClick={() => setPage((p) => p + 1)}
            disabled={page === totalPages - 1}
            className="text-sm text-gray-500 hover:text-gray-700 disabled:opacity-30 transition-colors"
          >
            次へ →
          </button>
        </div>
      )}
    </div>
  );
}
