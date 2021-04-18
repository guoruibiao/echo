package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoruibiao/echo/routers"
	"log"
)

func init() {
	// do
}

func main() {
    router := gin.Default()
	// 静态文件路由设置
	router.Static("/statics/", "./statics")
	router.Static("/templates/", "./templates")

    // 注册 html
    router.LoadHTMLGlob("./templates/*")

	// 动态注册服务路由
	routers.Init(router)

	// 服务启动
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
