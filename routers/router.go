package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/echo/controllers"
)
func Init(engine *gin.Engine) {
	indexRouter := engine.Group("/")
	{
		indexRouter.GET("/index", controllers.HomeIndex)
		indexRouter.POST("/index", controllers.HomeIndex)
	}

	proxyRouter := engine.Group("/")
	{
		proxyRouter.Any("/proxy", controllers.Proxy)
	}
}
