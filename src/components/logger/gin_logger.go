package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinMiddleLogger() gin.HandlerFunc {

	return func(context *gin.Context) {

		// 开始时间
		start := time.Now()
		// 处理请求
		context.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := context.Request.URL.Path

		clientIP := context.ClientIP()
		method := context.Request.Method
		statusCode := context.Writer.Status()

		latencyStr := fmt.Sprintf("%v", latency)

		errorMsg := context.Errors.ByType(gin.ErrorTypePrivate).String()

		GinLogger.Info("gin路由日志", zap.Int("statusCode", statusCode),
			zap.String("latency", latencyStr),
			zap.String("clientIP", clientIP),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("errorMsg", errorMsg))

	}

}
