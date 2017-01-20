package main

import (
	"GoReadNote/handlers"
	"GoReadNote/logger"
	"GoReadNote/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode  gin.DebugMode
	router := gin.Default()      //获得路由实例

	//添加中间件
	router.Use(middleware.Middleware)
	//搜索小说
	router.GET("/", handlers.HomeHandler)
	router.GET("/SearchNote", handlers.SearchNoteHandler)
	router.GET("GetBookInfo", handlers.GetBookInfoHandler)
	//获取章节内容
	router.GET("/GetBookContent", handlers.GetNoteContentHandler)

	router.GET("/GetJson", handlers.GetJsonHandler)

	logger.ALogger().Notice("Listen start.")
	//监听端口
	http.ListenAndServe(":8005", router)
}
