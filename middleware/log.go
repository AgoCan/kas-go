package middleware

/*
# 日志文件默认存储位置
 `/var/logging/das-go`
使用配置文件进行配置

https://github.com/sirupsen/logrus

*/

import (
	"kas-go/utils/logging"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"kas-go/config"
)

// LogMiddleware 初始化
func LogMiddleware() gin.HandlerFunc {
	// 日志对应yaml配置文件logAutoFile
	logging.Init(config.LogAutoFile)
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		Method := c.Request.Method

		// 请求路由
		RequestURI := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logging.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"method":       Method,
			"request_uri":  RequestURI,
		}).Info()

	}

}
