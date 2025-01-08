package controller

import (
	"github.com/gin-gonic/gin"
	"voice-assistant-manager/controller/assistant"
	"voice-assistant-manager/controller/voices"
)

// 实例化router结构体，可以使用该对象点出首字母大写的方法（跨包调用）
var Router router

// 声明一个router结构体
type router struct{}

// 初始化路由规则
func (r *router) InitApiRouter(router *gin.Engine) {
	v1Party := router.Group("/api")
	assistant.Install(v1Party)
	voices.Install(v1Party)
}
