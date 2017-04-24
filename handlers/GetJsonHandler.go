package handlers

import (
	"GoReadNovel/logger"
	"GoReadNovel/noveldb"
	"GoReadNovel/spider"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

//返回Json的一个模板  Code在不同情况下有不同作用
type JsonRet struct {
	Code int         `json:"code"`
	Ret  interface{} `json:"ret"`
}

//下面是个基本的例子
func GetJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetJsonHandler")
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	holder := JsonHolder{Id: 1, Name: "123"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(200, holder)

	return
}

//旧，新的是查数据库的方式查novelname或者作者
/*
func GetSearchNovelJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNovelJsonHandler")

	h := gin.H{}

	novelName, exist := c.GetQuery("novelname")
	if !exist {
		c.JSON(500, h)
		return
	}
	novelListMap, find := spider.SearchNovelByName(novelName)
	if !find {
		//没有找到 再试试直接Get
		chptMap, ok := spider.GetNovelChapterListByNovelName(novelName)
		if !ok {
			c.JSON(500, h)
			return
		}
		var cpInfo []spider.ChapterInfo

		for i := 1; i <= len(chptMap); i++ {
			cpInfo = append(cpInfo, chptMap[i])
		}
		//code = 0 为一个结果  code = 1为小说列表
		retJson := JsonRet{Code: 0, Ret: cpInfo}
		c.JSON(200, retJson)
		return
	}
	var novelInfo []spider.SearchNovel

	for i := 1; i <= len(novelListMap); i++ {
		novelInfo = append(novelInfo, novelListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return
}
*/

//获取小说的文章内容
func GetNovelContentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNovelJsonHandler")

	url, exist := c.GetQuery("go")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "input err"}
		c.JSON(500, errJson)
		return
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chp := spider.GetNovelContent(url)
	if chp == nil {
		errJson := JsonRet{Code: 0, Ret: "spider err"}
		c.JSON(500, errJson)
		return
	}

	c.JSON(200, chp)
	return

}

/*
func GetTopNovelListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopNovelListJsonHandler")
	h := gin.H{}
	type JsonRet struct {
		Code int         `json:"code"` //code = 0 为一个结果  code = 1为小说列表
		Ret interface{} `json:"ret"`
	}
	novelListMap, ok := spider.GetTopNovelList()
	if !ok {
		c.JSON(500, h)
		return
	}
	var novelInfo []spider.TopNovel

	for i := 1; i <= len(novelListMap); i++ {
		novelInfo = append(novelInfo, novelListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return

}
*/

//获取小说内容通过类型以及页数 这个先留着
func GetTopByTypeNovelListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopByTypeNovelListJsonHandler")
	h := gin.H{}
	type JsonRet struct {
		Code int         `json:"code"` //code = 0 为一个结果  code = 1为小说列表
		Ret  interface{} `json:"ret"`
	}
	//获取
	novelType, exist := c.GetQuery("ntype")
	if !exist {
		novelType = "quanbu"
	}
	sortType, exist := c.GetQuery("stype")
	if !exist {
		sortType = "default"
	}
	page, exist := c.GetQuery("page")
	if !exist {
		page = "1"
	}

	novelListMap, ok := spider.GetTopByTypeNovelList(novelType, sortType, page)
	if !ok {
		c.JSON(500, h)
		return
	}
	var novelInfo []spider.TopTypeNovel

	for i := 1; i <= len(novelListMap); i++ {
		novelInfo = append(novelInfo, novelListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return
}

/*
func GetNovelInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNovelInfoJsonHandler")
	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chptMap, ok := spider.GetNovelChapterListByUrl(url)
	if !ok {
		c.JSON(500, h)
		return
	}
	var novelInfo []spider.ChapterInfo

	for i := 1; i <= len(chptMap); i++ {
		novelInfo = append(novelInfo, chptMap[i])
	}
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return

}
*/
type ListFiles struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

func GetFileListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetFileListJsonHandler")
	h := gin.H{}
	//filedir  main wei
	ftype, exist := c.GetQuery("filedir")
	logger.ALogger().Debugf("filedir = %s", ftype)
	if !exist {
		c.JSON(500, h)
		logger.ALogger().Error("没有发现filedir")
		return
	}
	var dir string
	if ftype == "Main" {
		dir = Upload_Dir + "main/"
	} else if ftype == "Wei" {
		dir = Upload_Dir + "wei/"
	}

	lm := make([]ListFiles, 0)
	//遍历目录，读出文件名称 大小
	filepath.Walk(dir, func(path string, fi os.FileInfo, err error) error {
		if nil == fi {
			return err
		}
		if fi.IsDir() {
			return nil
		}
		var m ListFiles
		m.Name = fi.Name()
		m.Size = strconv.FormatInt(fi.Size()/1024, 10)
		lm = append(lm, m)
		return nil
	})

	c.JSON(200, lm)
}

//重写
//搜索小说通过小说名称或者作者 重写
func GetSearchNovelByNameOrAuthorJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNovelByNameOrAuthorJsonHandler Re")

	//get的data变成了key
	key, exist := c.GetQuery("key")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "input err"}
		c.JSON(500, errJson)
		return
	}
	novelListmap, find := noveldb.FindDatasFromNovelByNameOrAuthor(key)
	if !find {
		//数据库中没有找到，之后可以添加使用爬虫爬取一下
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return
	}

	var novelInfo []noveldb.Novel
	for i := 1; i <= len(novelListmap); i++ {
		novelInfo = append(novelInfo, novelListmap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return
}

//获取小说info  重写 前端应该加入相应页面
func GetNovelInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNovelInfoJsonHandler Re")
	//这个id是小说在数据库中的ID
	id, exist := c.GetQuery("id")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "input err"}
		c.JSON(500, errJson)
		return
	}
	novelId, err := strconv.Atoi(id)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "input err"}
		c.JSON(500, errJson)
		return
	}
	novelInfo, find := noveldb.FindOneDataFromNovelByID(novelId)
	if !find {
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return
	}

	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return

}

//获取小说的Top50 有一个type的
func GetTopNovelListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopNovelListJsonHandler Re")
	//应该有一个type的 没有值得话默认排序 allvote, goodnum 两种类型
	topty, exist := c.GetQuery("toptype")
	var novelListMap map[int]noveldb.Novel
	if !exist {
		var find bool
		novelListMap, find = noveldb.FindDatasFromNovel(0, 50)
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
	}
	if topty == "allvote" {
		novelList, find := noveldb.FindDatasFromAllVote(0, 50)
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
		for i := 0; i < len(novelList); i++ {
			novel, _ := noveldb.FindOneDataFromNovelByID(novelList[i].NovelID)
			novelListMap[i] = novel
		}
	} else if topty == "goodnum" {
		novelList, find := noveldb.FindDatasFromGoodNum(0, 50)
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
		for i := 0; i < len(novelList); i++ {
			novel, _ := noveldb.FindOneDataFromNovelByID(novelList[i].NovelID)
			novelListMap[i] = novel
		}
	}

	var novelsInfo []noveldb.Novel

	for i := 0; i < len(novelListMap); i++ {
		novelsInfo = append(novelsInfo, novelListMap[i])
	}
	// code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelsInfo}
	c.JSON(200, retJson)
	return
}

//获取小说list  这个貌似稍作修改即可
func GetChapterListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetChapterListJsonHandler Re")
	//这个id是小说在数据库中的ID
	url, exist := c.GetQuery("url")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "input err"}
		c.JSON(500, errJson)
		return
	}

	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chptMap, ok := spider.GetNovelChapterListByUrl(url)
	if !ok {
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return
	}
	var novelInfo []spider.ChapterInfo

	for i := 1; i <= len(chptMap); i++ {
		novelInfo = append(novelInfo, chptMap[i])
	}
	retJson := JsonRet{Code: 1, Ret: novelInfo}
	c.JSON(200, retJson)
	return
}

//获取对应类型的小说若干数量 新增
func GetATypeNovelJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetChapterListJsonHandler Re")
	novelType, exist := c.GetQuery("noveltype")
	if !exist {
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return

	}
	//noveltype转换查询数据库
	logger.ALogger().Debugf("Find Novel Type In DB :%s", noveldb.NovelTypeEtoC[novelType])
	novelListMap, find := noveldb.FindDatasFromNovelByNovelType(noveldb.NovelTypeEtoC[novelType])
	if !find {
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return
	}

	var novelsInfo []noveldb.Novel
	for i := 1; i <= len(novelListMap); i++ {
		novelsInfo = append(novelsInfo, novelListMap[i])
	}
	// code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelsInfo}
	c.JSON(200, retJson)
	return

}

//加入书架 新增  在写完用户登录以及维护session之后
