// フレーズ
export interface Phrase {
  id: string;
  // 例: 今日は天気がいい
  japaneseText: string;
  // 例: It's a nice day today.
  englishText: string;
  order: number;
  createdAt: Date;
}

// セッション
export interface Session {
  id: string;
  name: string;
  createdAt: Date;
  updatedAt: Date;
  phrases?: Phrase[];
}
