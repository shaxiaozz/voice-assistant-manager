package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
	"voice-assistant-manager/controller"
	"voice-assistant-manager/global"
	"voice-assistant-manager/middle"
)

// 鉴权账号
var user = gin.Accounts{
	global.AuthUsername: global.AuthPassword,
}

func GinInit() {
	r := gin.Default()
	// 自定义Gin Log
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 自定义日志输出格式
		return fmt.Sprintf("{\"client_ip\":\"%s\",\"access_time\":\"%s\",\"http_method\":\"%s\",\"uri\":\"%s\",\"http_proto\":\"%s\",\"http_code\":%d,\"respone_time\":\"%s\",\"http_useragent\":\"%s\",\"error_msg\":\"%s\"}\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}), gin.Recovery())
	// 加载跨域中间件
	r.Use(middle.Cors(), gin.BasicAuth(user))
	// 跨包调用router的初始化方法
	controller.Router.InitApiRouter(r)
	// 启动gin server
	r.Run("0.0.0.0:9090")
}

func InitConfig() {
	if os.Getenv("CARTESIA_API_KEY") == "" {
		log.Fatalf("CARTESIA_API_KEY not set.....")
	} else {
		global.CartesiaApiKey = os.Getenv("CARTESIA_API_KEY")
	}

	// 初始化 Docker 客户端
	//err := docker.InitDocker()
	//if err != nil {
	//	log.Fatalf("初始化 Docker 客户端失败: %v", err)
	//}
}
