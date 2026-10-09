"use client";

import { useState, FormEvent } from "react";
import { JapaneseInputForm } from "@/components/JapaneseInputForm";
import { TranslationCandidates } from "@/components/TranslationCandidates";
import { SessionList } from "@/components/SessionList";
import { useTranslate } from "@/hooks/useTranslate";
import { useSession } from "@/hooks/useSession";

export default function Home() {
  const { candidates, isLoading: isTranslating, error, translate, clearCandidates } = useTranslate();
  const { user, isAuthorized, sessions, activeSessionId, isLoading: isSessionLoading, signIn, signOut, createNewSession, setActiveSession, addPhraseToSession, mergePhrasesInSession } = useSession();
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

  if (isSessionLoading && !user) {
    return (
      <main className="min-h-screen bg-gray-50 flex items-center justify-center">
        <p className="text-sm text-gray-400">読み込み中...</p>
      </main>
    );
  }

  if (!user) {
    return (
      <main className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-gray-900 mb-2">Chill English</h1>
          <p className="text-sm text-gray-500 mb-8">シャドーイング練習アプリ</p>
          <button
            onClick={signIn}
            className="flex items-center gap-3 mx-auto px-6 py-3 rounded-xl border border-gray-200 bg-white
              text-sm font-medium text-gray-700 hover:bg-gray-50 shadow-sm transition-colors"
          >
            <svg viewBox="0 0 24 24" className="w-5 h-5" xmlns="http://www.w3.org/2000/svg">
              <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
              <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
              <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l3.66-2.84z" fill="#FBBC05"/>
              <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
            </svg>
            Google でログイン
          </button>
        </div>
      </main>
    );
  }

  if (!isAuthorized) {
    return (
      <main className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <p className="text-sm text-gray-600 mb-4">{user.email} はアクセス権がありません。</p>
          <button
            onClick={signOut}
            className="text-sm text-red-500 hover:text-red-600 underline"
          >
            ログアウト
          </button>
        </div>
      </main>
    );
  }

  return (
    <main className="min-h-screen bg-gray-50">
      <div className="max-w-lg mx-auto px-4 py-8">
        {/* Header */}
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-2xl font-bold text-gray-900">Chill English</h1>
            <p className="text-sm text-gray-500 mt-1">シャドーイング練習アプリ</p>
          </div>
          <button
            onClick={signOut}
            className="text-xs text-gray-400 hover:text-gray-600 transition-colors"
          >
            ログアウト
          </button>
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
