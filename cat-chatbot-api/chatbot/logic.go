package chatbot

import "strings"

type ChatResponse struct {
	Reply string `json:"reply"`
	Action string `json:"action,omitempty"`
}

func GetBotReply(text string) ChatResponse {
	t := strings.ToLower(text)

	if strings.Contains(t, "dance") || strings.Contains(t, "เต้น" ){
		return ChatResponse{Reply: "เมี้ยว~ ดูฉันเต้นสิ! 💃🐾🎵", Action: "dance"}
	}
	if strings.Contains(t, "สวัสดี") || strings.Contains(t, "hello") || strings.Contains(t, "hi") {
        return ChatResponse{Reply: "สวัสดี~ เมี้ยว! ยินดีที่ได้รู้จักนะ! 🐱"}
    }
    if strings.Contains(t, "เสื้อ") || strings.Contains(t, "outfit") {
        return ChatResponse{Reply: "โอ้โห~ มีเสื้อสวยเยอะเลย! ลองดู Meow Size ได้นะ 👕"}
    }
    if strings.Contains(t, "ราคา") || strings.Contains(t, "price") {
        return ChatResponse{Reply: "ราคาเสื้อเริ่มต้น ฿259 นะ! 🏷️"}
    }
    if strings.Contains(t, "happy") || strings.Contains(t, "สุข") {
        return ChatResponse{Reply: "เย้~ Maffin ดีใจด้วยนะ! 😸🎉"}
    }
	
	return ChatResponse{Reply: "เมี้ยว~ ได้รับแล้ว! 🐾 ถามเรื่องเสื้อแมวได้นะ!"}
}