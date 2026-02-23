import { NextRequest, NextResponse } from "next/server";
import { GoogleGenerativeAI } from "@google/generative-ai";

const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY!);

export async function POST(req: NextRequest) {
  try {
    const { japaneseText } = await req.json();

    if (!japaneseText || typeof japaneseText !== "string") {
      return NextResponse.json(
        { error: "japaneseText is required" },
        { status: 400 }
      );
    }

    const model = genAI.getGenerativeModel({ model: "gemini-2.5-flash" });

    const prompt = `
You are an English translation assistant for Japanese learners practicing shadowing.
Translate the following Japanese text into a single English sentence that sounds completely natural in everyday conversation — the kind a native speaker would actually say to a friend, not a textbook phrase.
Prioritize natural spoken rhythm, common contractions, and colloquial word choice over literal accuracy.
Return ONLY a JSON array containing exactly 1 string, no explanation, no markdown code blocks.

Japanese text: "${japaneseText}"

Example output format:
["translation"]
`;

    const result = await model.generateContent(prompt);
    const text = result.response.text();

    // Extract JSON array, handling possible markdown code blocks
    const match = text.match(/\[[\s\S]*\]/);
    if (!match) {
      return NextResponse.json(
        { error: "Failed to parse translation response" },
        { status: 500 }
      );
    }

    const candidates: string[] = JSON.parse(match[0]);

    return NextResponse.json({ candidates });
  } catch (error) {
    console.error("Translation error:", error);
    return NextResponse.json(
      { error: "Translation failed" },
      { status: 500 }
    );
  }
}
