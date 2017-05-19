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

	router.LoadHTMLGlob("templates/*.html")
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
	//获取首页url图片
	router.GET("/GetIndexImgsJson", handlers.GetIndexImgsJsonHandler)
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
	//增加书评
	router.GET("/AddANovelCommentJson", handlers.AddANovelCommentJsonHandler)
	//更新点赞数
	router.GET("/UpdateANovelCommentJson", handlers.UpdateANovelCommentJsonHandler)
	//获取书评
	router.GET("/GetANovelCommentsJson", handlers.GetANovelCommentsJsonHandler)
	//删除书评
	router.GET("/DeleteANovelCommentJson", handlers.DeleteANovelCommentJsonHandler)
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

	//table测试
	router.GET("/table", handlers.TableHandler)
	router.StaticFS("/Content", http.Dir("./webpage/Content"))
	router.StaticFS("/Scripts", http.Dir("./webpage/Scripts"))

	//----------------------Admin Begin------------------------
	//获取书籍Table信息
	router.GET("/GetNovelTableInfoJson", handlers.GetNovelTableInfoJsonHandler)
	//编辑table中的某本书籍
	router.GET("/GetEditNovelJson", handlers.GetEditNovelJsonHandler)
	//增加一本新的书籍
	router.GET("/GetANewNovelJson", handlers.GetANewNovelJsonHandler)
	//删除某本书籍
	router.GET("/GetDeleteNovleID", handlers.GetDeleteNovleIDHandler)
	//获取书籍数量
	router.GET("/GetNovelsCount", handlers.GetNovelsCountHandler)
	//终极模糊搜索小说可以通过 id、名字、作者 、类型获取书籍
	router.GET("/GetUltimateSearchNovelsJson", handlers.GetUltimateSearchNovelsJsonHandler)
	//获取用户信息
	router.GET("/GetUsersInfoJson", handlers.GetUsersInfoJsonHandler)
	//获取服务器端SpiderCofig信息
	router.GET("/GetSpiderConfigJson", handlers.GetSpiderConfigJsonHandler)
	//测试修改的正则表达式是否正确
	router.GET("/TestConfigJson", handlers.TestConfigJsonHandler)
	//保存修改的正则表达式
	router.GET("/SaveConfigJson", handlers.SaveConfigJsonHandler)
	//----------------------Admin End---------------------

	logger.ALogger().Notice("Listen start.")
	logger.ALogger().Notice("Listen 443 https")
	//监听端口
	//err := http.ListenAndServe(":8000", router)
	//http.ListenAndServeTLS(":443", "server.crt", "server.key", router)
	//8000端口是测试之用 实际端口为443
	err := http.ListenAndServeTLS(":443", "./ca/1_fsnsaber.cn_bundle.crt", "./ca/2_fsnsaber.cn.key", router)
	logger.ALogger().Error(err)

}
