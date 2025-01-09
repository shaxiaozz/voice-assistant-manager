package assistant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"voice-assistant-manager/service/assistant"
)

type Handler struct {
	assistantService assistant.Service
}

func NewHandler() *Handler {
	return &Handler{
		assistantService: assistant.NewService(),
	}
}

// 获取助手数据
func (h *Handler) List(ctx *gin.Context) {
	// 获取数据
	data, err := h.assistantService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1006,
			"msg":  "获取助手数据失败" + err.Error(),
		})
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1000,
		"msg":  "获取助手数据成功",
		"data": data,
	})
}

// 修改助手语音ID
func (h *Handler) UpdateVoiceID(ctx *gin.Context) {
	params := new(struct {
		VoiceID string `json:"voice_id"`
	})
	if err := ctx.ShouldBindJSON(params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": 1006,
			"msg":  "参数错误" + err.Error(),
		})
		return
	}

	// 修改ID
	if err := h.assistantService.UpdateVoiceID(params.VoiceID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 1006,
			"msg":  "修改助手语音ID失败" + err.Error(),
		})
		return
	}

	// 返回数据
	ctx.JSON(http.StatusOK, gin.H{
		"code":      1000,
		"msg":       "修改助手语音ID成功",
		"now_voice": params.VoiceID,
	})
}

func Install(parent *gin.RouterGroup) {
	handler := NewHandler()
	sp := parent.Group("/voice-assistant")
	sp.GET("/list", handler.List)
	sp.PUT("/update-voice-id", handler.UpdateVoiceID)
}
