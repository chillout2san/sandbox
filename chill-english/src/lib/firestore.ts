import {
  collection,
  doc,
  addDoc,
  setDoc,
  getDocs,
  serverTimestamp,
  orderBy,
  query,
  where,
  writeBatch,
  Timestamp,
} from "firebase/firestore";
import { db } from "./firebase";
import { Phrase, Session } from "./types";

export async function createSession(userId: string, name: string): Promise<string> {
  const ref = await addDoc(collection(db, "sessions"), {
    userId,
    name,
    createdAt: serverTimestamp(),
    updatedAt: serverTimestamp(),
  });
  return ref.id;
}

export async function addPhrase(
  sessionId: string,
  japaneseText: string,
  englishText: string,
  order: number
): Promise<string> {
  const phrasesRef = collection(db, "sessions", sessionId, "phrases");
  const ref = await addDoc(phrasesRef, {
    japaneseText,
    englishText,
    order,
    createdAt: serverTimestamp(),
  });

  // update session updatedAt
  await setDoc(
    doc(db, "sessions", sessionId),
    { updatedAt: serverTimestamp() },
    { merge: true }
  );

  return ref.id;
}

export async function getPhrases(sessionId: string): Promise<Phrase[]> {
  const phrasesRef = collection(db, "sessions", sessionId, "phrases");
  const q = query(phrasesRef, orderBy("order", "asc"));
  const snapshot = await getDocs(q);
  return snapshot.docs.map((d) => {
    const data = d.data();
    return {
      id: d.id,
      japaneseText: data.japaneseText,
      englishText: data.englishText,
      order: data.order,
      createdAt: (data.createdAt as Timestamp)?.toDate() ?? new Date(),
    };
  });
}

export async function mergePhrases(sessionId: string, phrases: Phrase[]): Promise<string> {
  const batch = writeBatch(db);
  const phrasesRef = collection(db, "sessions", sessionId, "phrases");

  // 既存フレーズを全削除
  for (const phrase of phrases) {
    batch.delete(doc(phrasesRef, phrase.id));
  }
  await batch.commit();

  // 結合フレーズを1件追加
  const japaneseText = phrases.map((p) => p.japaneseText).join("\n");
  const englishText = phrases.map((p) => p.englishText).join(" ");
  const ref = await addDoc(phrasesRef, {
    japaneseText,
    englishText,
    order: 0,
    createdAt: serverTimestamp(),
  });

  await setDoc(
    doc(db, "sessions", sessionId),
    { updatedAt: serverTimestamp() },
    { merge: true }
  );

  return ref.id;
}

export async function getSessions(userId: string): Promise<Session[]> {
  const q = query(collection(db, "sessions"), where("userId", "==", userId));
  const snapshot = await getDocs(q);
  const sessions = await Promise.all(
    snapshot.docs.map(async (d) => {
      const data = d.data();
      const phrases = await getPhrases(d.id);
      return {
        id: d.id,
        name: data.name ?? "無題",
        createdAt: (data.createdAt as Timestamp)?.toDate() ?? new Date(),
        updatedAt: (data.updatedAt as Timestamp)?.toDate() ?? new Date(),
        phrases,
      };
    })
  );
  return sessions.sort((a, b) => b.createdAt.getTime() - a.createdAt.getTime());
}
