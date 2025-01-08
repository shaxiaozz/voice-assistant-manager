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

func Install(parent *gin.RouterGroup) {
	handler := NewHandler()
	sp := parent.Group("/voice-assistant")
	sp.GET("/list", handler.List)
}
