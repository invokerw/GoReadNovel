package main

import (
	"GoReadNovel/handlers"
	"GoReadNovel/logger"
	"GoReadNovel/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	gin.SetMode(gin.ReleaseMode) //全局设置环境，此为开发环境，线上环境为 gin.ReleaseMode  gin.DebugMode
	router := gin.Default()      //获得路由实例

	router.LoadHTMLGlob("templates/*")
	//网页请求----------------------Begin
	//添加中间件
	router.Use(middleware.Middleware)
	//搜索小说
	router.GET("/", handlers.HomeHandler)
	router.GET("/SearchIndex", handlers.GetSearchIndexHandler)
	router.GET("/SearchNovel", handlers.SearchNovelHandler)
	router.GET("/GetBookInfo", handlers.GetBookInfoHandler)
	//获取章节内容
	router.GET("/GetBookContent", handlers.GetNovelContentHandler)

	//网页请求---------------------End

	//微信小程序登录
	router.GET("/WxOnLogin", handlers.WeiXinOnLoginHandler)

	//微信小程序JSON请求----------------------Begin
	//测试
	router.GET("/GetJson", handlers.GetJsonHandler)
	//搜索小说 小说名称或者作者  key = string
	router.GET("/GetSearchNovelJson", handlers.GetSearchNovelByNameOrAuthorJsonHandler)
	//获取小说内容   go = url
	router.GET("/GetNovelContentJson", handlers.GetNovelContentJsonHandler)
	//获取top榜上的小说  默认为总点击  toptype = allvote , goodnum
	router.GET("/GetTopNovelListJson", handlers.GetTopNovelListJsonHandler)
	//获取一种类型的小说 noveltype  = 汉语拼音
	router.GET("/GetATypeNovelJson", handlers.GetATypeNovelJsonHandler)
	//获取小说详细信息 id  = 数据库小说id
	router.GET("/GetNovelInfoJson", handlers.GetNovelInfoJsonHandler)
	//获取小说章节列表 url  = 小说地址
	router.GET("/GetChpListJson", handlers.GetChapterListJsonHandler)
	//这个暂时没有用到 与 GetTopNovelListJson  可以获取各种类型的各种小说 可以看 相应handler
	router.GET("/GetTopByTypeNovelListJson", handlers.GetTopByTypeNovelListJsonHandler)
	//增加小说到书架
	router.GET("/AddAUserNovelInBookShelfJson", handlers.AddAUserNovelInBookShelfJsonHandler)
	//更新读到的章节
	router.GET("/UpdateAUserNovelInBookShelfJson", handlers.UpdateAUserNovelInBookShelfJsonHandler)
	//删除书架中的某个小说 user session  novel id
	router.GET("/DeleteAUserNovelInBookShelfJson", handlers.DeleteAUserNovelInBookShelfJsonHandler)
	//获取书架上的所有书籍数据 user session
	router.GET("/GetUserBookShelfNovelsJson", handlers.GetUserBookShelfNovelsJsonHandler)
	//这本书是否在书架上 user session  novel id
	router.GET("/GetTheNovelInBookShelfJson", handlers.GetTheNovelInBookShelfJsonHandler)
	//微信小程序JSON请求----------------------End

	//文件上传
	router.GET("/UploadFile", handlers.GetUpLoadPageHandler)
	router.POST("/UploadFile", handlers.UploadFileHandler)
	router.GET("/GetFileListJson", handlers.GetFileListJsonHandler)
	router.StaticFS("/Main", http.Dir("./savefile/main"))
	router.StaticFS("/Weifei", http.Dir("./savefile/wei"))
	//icon
	router.StaticFile("/favicon.ico", "./statics/favicon.ico")

	//测试
	router.GET("/1", handlers.NewHomeHandler)
	router.StaticFS("/statics", http.Dir("./statics"))

	logger.ALogger().Notice("Listen start.")
	logger.ALogger().Notice("Listen 443 https")
	//监听端口
	//err := http.ListenAndServe(":8000", router)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	//8000端口是测试之用 实际端口为443
	err := http.ListenAndServeTLS(":443", "./ca/1_fsnsaber.cn_bundle.crt", "./ca/2_fsnsaber.cn.key", router)
	logger.ALogger().Error(err)

}
