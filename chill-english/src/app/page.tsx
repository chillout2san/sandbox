"use client";

import { useState, FormEvent } from "react";
import { JapaneseInputForm } from "@/components/JapaneseInputForm";
import { TranslationCandidates } from "@/components/TranslationCandidates";
import { SessionList } from "@/components/SessionList";
import { useTranslate } from "@/hooks/useTranslate";
import { useSession } from "@/hooks/useSession";

export default function Home() {
  const { candidates, isLoading: isTranslating, error, translate, clearCandidates } = useTranslate();
  const { sessions, activeSessionId, isLoading: isSessionLoading, createNewSession, setActiveSession, addPhraseToSession, mergePhrasesInSession } = useSession();
  const [currentJapanese, setCurrentJapanese] = useState("");
  const [playingId, setPlayingId] = useState<string | null>(null);
  const [newSessionName, setNewSessionName] = useState("");
  const [showNewSessionForm, setShowNewSessionForm] = useState(false);

  const handleTranslate = (text: string) => {
    setCurrentJapanese(text);
    translate(text);
  };

  const handleSelectCandidate = async (englishText: string) => {
    await addPhraseToSession(currentJapanese, englishText);
    clearCandidates();
    setCurrentJapanese("");
  };

  const handleCreateSession = async (e: FormEvent) => {
    e.preventDefault();
    const name = newSessionName.trim();
    if (!name) return;
    await createNewSession(name);
    setNewSessionName("");
    setShowNewSessionForm(false);
  };

  const handlePlay = async (text: string, id: string) => {
    if (playingId) return;
    setPlayingId(id);
    try {
      const res = await fetch("/api/tts", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ text }),
      });

      if (!res.ok) throw new Error("TTS request failed");

      const { audioContent } = await res.json();
      const binary = atob(audioContent);
      const bytes = new Uint8Array(binary.length).map((_, i) =>
        binary.charCodeAt(i)
      );
      const url = URL.createObjectURL(new Blob([bytes], { type: "audio/mp3" }));
      const audio = new Audio(url);
      audio.onended = () => {
        URL.revokeObjectURL(url);
        setPlayingId(null);
      };
      audio.onerror = () => {
        URL.revokeObjectURL(url);
        setPlayingId(null);
      };
      await audio.play();
    } catch (err) {
      console.error("Play error:", err);
      setPlayingId(null);
    }
  };

  return (
    <main className="min-h-screen bg-gray-50">
      <div className="max-w-lg mx-auto px-4 py-8">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-2xl font-bold text-gray-900">Chill English</h1>
          <p className="text-sm text-gray-500 mt-1">シャドーイング練習アプリ</p>
        </div>

        {/* Input Section */}
        <div className="bg-white rounded-2xl shadow-sm border border-gray-100 p-5 mb-6">
          {activeSessionId ? (
            <>
              <JapaneseInputForm onTranslate={handleTranslate} isLoading={isTranslating} />
              {error && <p className="mt-3 text-sm text-red-500">{error}</p>}
              <TranslationCandidates
                candidates={candidates}
                onSelect={handleSelectCandidate}
                onPlay={handlePlay}
                playingId={playingId}
              />
            </>
          ) : (
            <p className="text-sm text-gray-400 text-center py-4">
              セッションを選択するか、新規作成してください
            </p>
          )}
        </div>

        {/* Session List */}
        <div>
          <div className="flex items-center justify-between mb-3">
            <h2 className="text-sm font-semibold text-gray-700">セッション</h2>
            {!showNewSessionForm && (
              <button
                onClick={() => setShowNewSessionForm(true)}
                className="text-xs text-emerald-600 hover:text-emerald-700 font-medium"
              >
                ＋ 新規セッション
              </button>
            )}
          </div>

          {showNewSessionForm && (
            <form onSubmit={handleCreateSession} className="flex gap-2 mb-4">
              <input
                type="text"
                value={newSessionName}
                onChange={(e) => setNewSessionName(e.target.value)}
                placeholder="セッション名を入力..."
                autoFocus
                className="flex-1 px-4 py-2 rounded-xl border border-gray-200 bg-white text-sm text-gray-900
                  focus:outline-none focus:ring-2 focus:ring-emerald-400 focus:border-transparent
                  placeholder-gray-400"
              />
              <button
                type="submit"
                disabled={!newSessionName.trim()}
                className="px-4 py-2 rounded-xl bg-emerald-500 text-white text-sm font-medium
                  hover:bg-emerald-600 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                作成
              </button>
              <button
                type="button"
                onClick={() => { setShowNewSessionForm(false); setNewSessionName(""); }}
                className="px-4 py-2 rounded-xl border border-gray-200 text-sm text-gray-600
                  hover:bg-gray-50 transition-colors"
              >
                キャンセル
              </button>
            </form>
          )}

          {isSessionLoading ? (
            <div className="text-center py-12 text-gray-400 text-sm">
              読み込み中...
            </div>
          ) : (
            <SessionList
              sessions={sessions}
              activeSessionId={activeSessionId}
              playingId={playingId}
              onSelectSession={setActiveSession}
              onPlay={handlePlay}
              onMergePhrases={mergePhrasesInSession}
            />
          )}
        </div>
      </div>
    </main>
  );
}
