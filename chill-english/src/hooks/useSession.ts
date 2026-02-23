"use client";

import { useState, useEffect, useCallback } from "react";
import { onAuthStateChanged, signInAnonymously } from "firebase/auth";
import { auth } from "@/lib/firebase";
import { Session, Phrase } from "@/lib/types";
import { createSession, addPhrase, getSessions } from "@/lib/firestore";

const ACTIVE_SESSION_KEY = "chill-english-active-session";

interface UseSessionResult {
  sessions: Session[];
  activeSessionId: string | null;
  isLoading: boolean;
  createNewSession: (name: string) => Promise<void>;
  setActiveSession: (id: string) => void;
  addPhraseToSession: (japaneseText: string, englishText: string) => Promise<void>;
}

async function ensureAuth(): Promise<string> {
  return new Promise((resolve, reject) => {
    const unsubscribe = onAuthStateChanged(auth, async (user) => {
      unsubscribe();
      if (user) {
        resolve(user.uid);
      } else {
        signInAnonymously(auth)
          .then(({ user: newUser }) => resolve(newUser.uid))
          .catch(reject);
      }
    });
  });
}

export function useSession(): UseSessionResult {
  const [sessions, setSessions] = useState<Session[]>([]);
  const [activeSessionId, setActiveSessionId] = useState<string | null>(null);
  const [userId, setUserId] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const init = async () => {
      setIsLoading(true);
      try {
        const uid = await ensureAuth();
        setUserId(uid);
        const allSessions = await getSessions(uid);
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
      } finally {
        setIsLoading(false);
      }
    };
    init();
  }, []);

  const createNewSession = useCallback(
    async (name: string) => {
      if (!userId) return;
      const id = await createSession(userId, name);
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
    [userId]
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

  return { sessions, activeSessionId, isLoading, createNewSession, setActiveSession, addPhraseToSession };
}
