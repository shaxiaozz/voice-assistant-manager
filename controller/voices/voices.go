package voices

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voice-assistant-manager/service/assistant"
	"voice-assistant-manager/service/voices"
)

type Handler struct {
	voicesService    voices.Service
	assistantService assistant.Service
}

func NewHandler() *Handler {
	return &Handler{
		voicesService:    voices.NewService(),
		assistantService: assistant.NewService(),
	}
}

// VoiceInfo 表示语音信息的结构体
type VoiceInfo struct {
	ID          string `json:"id"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Gender      string `json:"gender,omitempty"`
}

// ListCartesia 获取 Cartesia 语音列表
func (h *Handler) ListCartesia(ctx *gin.Context) {
	// 获取数据
	data, err := h.voicesService.ListCartesiaVoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1006,
			"msg":  "获取Cartesia语音列表失败: " + err.Error(),
		})
		return
	}

	// 转换数据格式
	voiceList := make([]VoiceInfo, 0, len(data))
	for _, voice := range data {
		if voice.Description == "" || voice.Gender == "" {
			continue
		}

		voiceList = append(voiceList, VoiceInfo{
			ID:          voice.Id,
			Language:    voice.Language,
			Description: voice.Description,
			Gender:      voice.Gender,
		})
	}

	// 获取现在助手的语音ID
	// 获取数据
	assistantData, err := h.assistantService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1006,
			"msg":  "获取Cartesia语音列表失败" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":      1000,
		"msg":       "获取Cartesia语音列表成功",
		"data":      voiceList,
		"now_voice": assistantData[0].Voice,
	})
}

func Install(parent *gin.RouterGroup) {
	handler := NewHandler()
	sp := parent.Group("/voice")
	sp.GET("/cartesia/list", handler.ListCartesia)
}
