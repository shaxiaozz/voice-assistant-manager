package global

import "time"

var (
	AuthUsername   = "voice-assistant-manager"
	AuthPassword   = "voice-assistant-manager"
	CartesiaApiKey string
)

const DefaultAssistant = `[
    {
        "id": "lawyer",
        "name": "律师助手",
        "title": "专业法律咨询",
        "description": "专业法律咨询师",
        "voice": "a998a552-22ac-4457-aada-b192ed6b2eaf"
    },
    {
        "id": "stewardess",
        "name": "空乘助手",
        "title": "航空旅行服务",
        "description": "航空旅行咨询",
        "voice": "a998a552-22ac-4457-aada-b192ed6b2eaf"
    },
    {
        "id": "xiaomei",
        "name": "贴心陪聊",
        "title": "小梅贴心陪聊助手",
        "description": "温暖治愈的AI聊天伙伴，擅长倾听和交流",
        "voice": "a998a552-22ac-4457-aada-b192ed6b2eaf"
    },
    {
        "id": "xiaonana",
        "name": "小娜娜女友",
        "title": "虚拟女友",
        "description": "甜美可爱的虚拟女友，让生活充满温暖与关怀",
        "voice": "a998a552-22ac-4457-aada-b192ed6b2eaf"
    },
    {
        "id": "psychologist",
        "name": "心理咨询师助手",
        "title": "林心怡心理咨询师",
        "description": "专业的AI心理咨询师，提供心理支持和倾听服务",
        "voice": "a998a552-22ac-4457-aada-b192ed6b2eaf"
    }
]`

// Assistant 定义数据结构
type Assistant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Voice       string `json:"voice"`
}

type CartesiaVoices struct {
	Id          string    `json:"id"`
	IsPublic    bool      `json:"is_public"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Embedding   []float64 `json:"embedding"`
	Language    string    `json:"language"`
	UserId      string    `json:"user_id"`
	BaseVoiceId string    `json:"base_voice_id"`
}
