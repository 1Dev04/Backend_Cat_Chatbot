package chatbot

import "strings"

type ChatResponse struct {
	Reply  string `json:"reply"`
	Action string `json:"action,omitempty"`
}

func GetBotReply(text string) ChatResponse {
    t := strings.ToLower(text)

    if strings.Contains(t, "dance") || strings.Contains(t, "เต้น") {
        return ChatResponse{
            Reply:  "เมี้ยว~ ดู Maffin โชว์สเต็ปหน่อยไหม! 💃🐾🎵",
            Action: "dance",
        }
    }

    if strings.Contains(t, "สวัสดี") || strings.Contains(t, "hello") || strings.Contains(t, "hi") {
        return ChatResponse{
            Reply: "สวัสดี~ เมี้ยว! ยินดีที่ได้รู้จักนะ 😸✨",
        }
    }

    if strings.Contains(t, "outfit") || strings.Contains(t, "ชุด") {
        return ChatResponse{
            Reply: "วันนี้จะใส่ชุดไหนดีน้า~ 👕 ลองไปดูในหน้า Search ได้เลย!",
        }
    }

    // ❗ เอา "hi" ออก ไม่งั้นชนกับ greeting
    if strings.Contains(t, "scan meow") || strings.Contains(t, "scan moew") || strings.Contains(t, "สแกนแมว") {
        return ChatResponse{
            Reply: "เมี้ยว~ เปิดโหมด Meow Size แล้ว! 📸 → 🤖 → 📊 วิเคราะห์ไซส์ให้เรียบร้อย!",
        }
    }

    if strings.Contains(t, "love") || strings.Contains(t, "รัก") {
        return ChatResponse{
            Reply:  "Maffin เขินนะเนี่ย~ 😸💕 ตกหลุมรักแล้ววว",
            Action: "happy",
        }
    }

    if strings.Contains(t, "8bit") || strings.Contains(t, "8 bit") || strings.Contains(t, "8-bit") {
        return ChatResponse{
            Reply:  "ติ๊ง! ✨ Maffin แปลงร่างเป็น 8-Bit แล้ว! 🎮😸 น่ารักมั้ยล่ะ~",
            Action: "8bit",
        }
    }

    if strings.Contains(t, "cmd") || strings.Contains(t, "command") || strings.Contains(t, "ช่วย") {
        return ChatResponse{
            Reply: "(Command)\n🐾 dance\n👕 outfit\n👋 hello\n🎮 8bit\n💕 love",
            Action: "command",
        }
    }

    return ChatResponse{
        Reply: "เมี้ยว~ ได้ยินแล้วนะ! 🐾 ลองถามเรื่องชุดหรือเล่นกับ Maffin ดูสิ 😏",
    }
}