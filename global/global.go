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
    "voice": "bafcab8d-d391-44fe-9711-e5c94e899f43"
  },
  {
    "id": "stewardess",
    "name": "空乘助手",
    "title": "航空旅行服务",
    "description": "航空旅行咨询",
    "voice": "bafcab8d-d391-44fe-9711-e5c94e899f43"
  }
]
`

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
	Gender      string    `json:"gender,omitempty"`
	Embedding   []float64 `json:"embedding"`
	Language    string    `json:"language"`
	ApiStatus   string    `json:"api_status"`
	UserId      string    `json:"user_id,omitempty"`
}
