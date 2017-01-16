package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode  gin.DebugMode
	router := gin.Default()      //获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/simple/server/get", handlers.GetHandler)
	router.POST("/simple/server/post", handlers.PostHandler)
	router.PUT("/simple/server/put", handlers.PutHandler)
	router.DELETE("/simple/server/delete", handlers.DeleteHandler)
	//监听端口
	http.ListenAndServe(":8005", router)
}
