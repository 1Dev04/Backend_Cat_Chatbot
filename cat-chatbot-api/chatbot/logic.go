package chatbot

import "strings"

type ChatResponse struct {
	Reply  string `json:"reply"`
	Action string `json:"action,omitempty"`
}

func GetBotReply(text string) ChatResponse {
    t := strings.ToLower(text)

    if strings.Contains(t, "dance") || strings.Contains(t, "เต้น") {
        return ChatResponse{Reply: "เมี้ยว~ ดูฉันเต้นสิ! 💃🐾🎵", Action: "dance"}
    }
    if strings.Contains(t, "สวัสดี") || strings.Contains(t, "hello") || strings.Contains(t, "hi") {
        return ChatResponse{Reply: "สวัสดี~ เมี้ยว! ยินดีที่ได้รู้จักนะ! 🐱"}
    }
    if strings.Contains(t, "outfit") || strings.Contains(t, "มีชุดอะไรบ้าง") {
        return ChatResponse{Reply: "โอ้โห~ มีเสื้อ outfit เยอะเลย! ลองดู Search ได้เลยนะ 👕"}
    }
    if strings.Contains(t, "love") || strings.Contains(t, "รัก") {
        return ChatResponse{Reply: "Maffin ตกหลุมรักแล้ววว 😸💕", Action: "happy"} 
    }
    if strings.Contains(t, "8bit") || strings.Contains(t, "8 bit") || strings.Contains(t, "8-bit") {
        return ChatResponse{Reply: "Maffin ร่าง 8 Bit! น่ารักไหม 🎮😸", Action: "8bit"}
    }

    return ChatResponse{Reply: "เมี้ยว~ ได้รับแล้ว! 🐾 ถามเรื่องเสื้อแมวได้นะ!"}
}
