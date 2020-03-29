package routers

import (
	"github.com/gin-gonic/gin"
	"kas-go/handler/kubernetes"

	"kas-go/middleware"
)

// SetupRouter 路由路口
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// docker 相关操作
	router.Use(middleware.LogMiddleware())
	dockerGroup := router.Group("/docker")
	{
		dockerGroup.POST("/run", kubernetes.PodHandler)

	}
	return router
}
