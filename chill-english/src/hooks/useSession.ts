"use client";

import { useState, useEffect, useCallback } from "react";
import {
  onAuthStateChanged,
  signInWithPopup,
  signOut as firebaseSignOut,
  User,
} from "firebase/auth";
import { auth, googleProvider } from "@/lib/firebase";
import { Session, Phrase } from "@/lib/types";
import { createSession, addPhrase, getSessions, mergePhrases } from "@/lib/firestore";

const ACTIVE_SESSION_KEY = "chill-english-active-session";
const ALLOWED_EMAIL = process.env.NEXT_PUBLIC_ALLOWED_EMAIL;

interface UseSessionResult {
  user: User | null;
  isAuthorized: boolean;
  sessions: Session[];
  activeSessionId: string | null;
  isLoading: boolean;
  signIn: () => Promise<void>;
  signOut: () => Promise<void>;
  createNewSession: (name: string) => Promise<void>;
  setActiveSession: (id: string) => void;
  addPhraseToSession: (japaneseText: string, englishText: string) => Promise<void>;
  mergePhrasesInSession: (sessionId: string) => Promise<void>;
}

export function useSession(): UseSessionResult {
  const [user, setUser] = useState<User | null>(null);
  const [sessions, setSessions] = useState<Session[]>([]);
  const [activeSessionId, setActiveSessionId] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  const isAuthorized = !!user && (!ALLOWED_EMAIL || user.email === ALLOWED_EMAIL);

  useEffect(() => {
    const unsubscribe = onAuthStateChanged(auth, async (u) => {
      setUser(u);
      const authorized = !!u && (!ALLOWED_EMAIL || u.email === ALLOWED_EMAIL);
      if (authorized) {
        setIsLoading(true);
        try {
          const allSessions = await getSessions(u!.uid);
          setSessions(allSessions);
          const stored = localStorage.getItem(ACTIVE_SESSION_KEY);
          if (stored && allSessions.some((s) => s.id === stored)) {
            setActiveSessionId(stored);
          } else if (allSessions.length > 0) {
            setActiveSessionId(allSessions[0].id);
            localStorage.setItem(ACTIVE_SESSION_KEY, allSessions[0].id);
          }
        } catch (err) {
          console.error("Session init error:", err);
        }
      } else {
        setSessions([]);
        setActiveSessionId(null);
      }
      setIsLoading(false);
    });
    return unsubscribe;
  }, []);

  const signIn = useCallback(async () => {
    await signInWithPopup(auth, googleProvider);
  }, []);

  const signOut = useCallback(async () => {
    await firebaseSignOut(auth);
    localStorage.removeItem(ACTIVE_SESSION_KEY);
  }, []);

  const createNewSession = useCallback(
    async (name: string) => {
      if (!user) return;
      const id = await createSession(user.uid, name);
      const newSession: Session = {
        id,
        name,
        createdAt: new Date(),
        updatedAt: new Date(),
        phrases: [],
      };
      setSessions((prev) => [newSession, ...prev]);
      setActiveSessionId(id);
      localStorage.setItem(ACTIVE_SESSION_KEY, id);
    },
    [user]
  );

  const setActiveSession = useCallback((id: string) => {
    setActiveSessionId(id);
    localStorage.setItem(ACTIVE_SESSION_KEY, id);
  }, []);

  const addPhraseToSession = useCallback(
    async (japaneseText: string, englishText: string) => {
      if (!activeSessionId) return;
      const currentPhrases =
        sessions.find((s) => s.id === activeSessionId)?.phrases ?? [];
      const order = currentPhrases.length;
      const id = await addPhrase(activeSessionId, japaneseText, englishText, order);
      const newPhrase: Phrase = {
        id,
        japaneseText,
        englishText,
        order,
        createdAt: new Date(),
      };
      setSessions((prev) =>
        prev.map((s) =>
          s.id === activeSessionId
            ? { ...s, phrases: [...(s.phrases ?? []), newPhrase] }
            : s
        )
      );
    },
    [activeSessionId, sessions]
  );

  const mergePhrasesInSession = useCallback(
    async (sessionId: string) => {
      const phrases = sessions.find((s) => s.id === sessionId)?.phrases ?? [];
      if (phrases.length < 2) return;
      const newId = await mergePhrases(sessionId, phrases);
      const merged: Phrase = {
        id: newId,
        japaneseText: phrases.map((p) => p.japaneseText).join("\n"),
        englishText: phrases.map((p) => p.englishText).join(" "),
        order: 0,
        createdAt: new Date(),
      };
      setSessions((prev) =>
        prev.map((s) =>
          s.id === sessionId ? { ...s, phrases: [merged] } : s
        )
      );
    },
    [sessions]
  );

  return {
    user,
    isAuthorized,
    sessions,
    activeSessionId,
    isLoading,
    signIn,
    signOut,
    createNewSession,
    setActiveSession,
    addPhraseToSession,
    mergePhrasesInSession,
  };
}
