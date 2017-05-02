package handlers

import (
	"GoReadNovel/logger"
	"GoReadNovel/myredis"
	"GoReadNovel/noveldb"
	"GoReadNovel/spider"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v4"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	//这里加上看看是否书籍在书架上做一个记录，如果他登录有效，书架上有这本小说
	session, exist := c.GetQuery("session")
	if !exist {
		//errJson := JsonRet{Code: -2, Ret: "can't find session"}
		//c.JSON(500, errJson)
		logger.ALogger().Error("session not find")
		return
	}
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		//errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		//c.JSON(500, errJson)
		logger.ALogger().Error("redis is nil")
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return
	}
	//如果redis中有这个key的话那就再给他续一段时间
	//redis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))
	datas := strings.Split(strings.TrimSpace(chp.Url), "/")
	if len(datas) != 7 {
		logger.ALogger().Debug("The Split Now Url datas = ", datas)
		return
	}
	novelUrl := spider.URL + "/book/" + datas[4] + "/" + datas[5] + "/"
	logger.ALogger().Debug("The Novel Url Is = ", novelUrl)
	if novel, find := noveldb.FindOneDataFromNovelByAddr(novelUrl); !find {
		logger.ALogger().Error("Not Find A Novel By The Url = ")
		return
	} else {
		if bookShelf, find := noveldb.FindOneNovelFromBookShelfByUserIDAndNovelID(uid, novel.ID); find {
			//发现了就更新一下bookShelf
			bookShelf.ReadChapterName = chp.ChapterName
			bookShelf.ReadChapterUrl = chp.Url[len(spider.URL):len(chp.Url)]
			noveldb.UpdateOneDataToBookShlefByUserIDAndNovelID(bookShelf)
		}
	}
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
	for i := 0; i < len(novelListmap); i++ {
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
		novelListMap = make(map[int]noveldb.Novel)
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
		novelListMap = make(map[int]noveldb.Novel)
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
	} else if topty == "default" {
		var find bool
		novelListMap, find = noveldb.FindDatasFromNovel(0, 50)
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
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
	//这个从spider模块走的从1开始
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
	retNovelNum := 50 //这里限制了返回小说的数量后可以修改
	if retNovelNum < len(novelListMap) {
		retNovelNum = len(novelListMap)
	}
	for i := 0; i < retNovelNum; i++ {
		novelsInfo = append(novelsInfo, novelListMap[i])
	}
	// code = 1为小说列表
	retJson := JsonRet{Code: 1, Ret: novelsInfo}
	c.JSON(200, retJson)
	return

}

//查询书架
func GetUserBookShelfNovelsJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetUserBookShelfNovelsJsonHandler ")
	session, exist := c.GetQuery("session")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find session"}
		c.JSON(500, errJson)
		return
	}
	logger.ALogger().Debug("Get session:", session)
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		c.JSON(500, errJson)
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return
	}
	//如果redis中有这个key的话那就再给他续一段时间
	//redis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))

	bookShelf, find := noveldb.FindOneUserBookShlefFromBookShelfByUserID(uid)
	if !find {
		errJson := JsonRet{Code: 0, Ret: "not have novels in bookshelf"}
		c.JSON(200, errJson)
		return
	}
	type AllInfo struct {
		Novel     noveldb.Novel     `json:"novel"`
		UserNovel noveldb.BookShelf `json:"usernovel"`
	}
	var allInfos []AllInfo
	for i := 0; i < len(bookShelf); i++ {
		allinfo := AllInfo{}
		novel, _ := noveldb.FindOneDataFromNovelByID(bookShelf[i].NovelID)
		allinfo.Novel = novel
		allinfo.UserNovel = bookShelf[i]
		//这里把userid做处理不返回客户端
		allinfo.UserNovel.UserID = "0.0"
		allInfos = append(allInfos, allinfo)
	}
	okJson := JsonRet{Code: 1, Ret: allInfos}
	c.JSON(200, okJson)
	return

}

//添加书籍到书架
func AddAUserNovelInBookShelfJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to AddAUserNovelInBookShelfJsonHandler ")
	session, exist := c.GetQuery("session")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find session"}
		c.JSON(500, errJson)
		return
	}
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		c.JSON(500, errJson)
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return

	}
	//如果redis中有这个key的话那就再给他续一段时间
	//myredis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))

	novelid, exist := c.GetQuery("novelid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		return
	}
	nid, err := strconv.Atoi(novelid)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "novelid err"}
		c.JSON(500, errJson)
		return
	}
	//加到这里可以节省查找的时间，下面还有一个就是add这个原子操作，下面代码太耗时间，不过加了携程可能会好一些？
	if _, find := noveldb.FindOneNovelFromBookShelfByUserIDAndNovelID(uid, nid); find {
		okJson := JsonRet{Code: 1, Ret: "bookshelf have the novel"}
		c.JSON(200, okJson)
		return
	}
	//FIXME:这里可能会很慢。所以可以考虑使用携程先将数据插入然后再更新第一章数据
	bookShelf := noveldb.BookShelf{}
	bookShelf.UserID = uid
	bookShelf.NovelID = nid
	//刚阅读肯定是第1章，还是直接插入第一章吧。。
	novel, _ := noveldb.FindOneDataFromNovelByID(nid)
	chptMap, ok := spider.GetNovelChapterListByUrl(novel.NovelUrl)
	if !ok {
		errJson := JsonRet{Code: 0, Ret: "can't get first chapter"}
		c.JSON(500, errJson)
		return
	}
	bookShelf.ReadChapterName = chptMap[1].ChapterName
	bookShelf.ReadChapterUrl = chptMap[1].Url //这里需要记得改一下
	if _, find := noveldb.FindOneNovelFromBookShelfByUserIDAndNovelID(uid, nid); !find {
		noveldb.InsertOneDataToBookShelf(bookShelf)
	}
	okJson := JsonRet{Code: 1, Ret: "insert to bookshelf ok"}
	c.JSON(200, okJson)
	return
}

//删除书架中的某个书籍
func DeleteAUserNovelInBookShelfJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to DeleteAUserNovelInBookShelfJsonHandler ")
	session, exist := c.GetQuery("session")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find session"}
		c.JSON(500, errJson)
		return
	}
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		c.JSON(500, errJson)
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return
	}
	//如果redis中有这个key的话那就再给他续一段时间
	//redis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))

	novelid, exist := c.GetQuery("novelid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		return
	}
	nid, err := strconv.Atoi(novelid)
	if err != nil {
		errJson := JsonRet{Code: 0, Ret: "novelid err"}
		c.JSON(500, errJson)
		return
	}

	noveldb.DeleteOneDataToBookShelfByUseridAndNovelid(uid, nid)
	okJson := JsonRet{Code: 1, Ret: "delete to bookshelf ok"}
	c.JSON(200, okJson)
}

//更新书架中的某个书籍
func UpdateAUserNovelInBookShelfJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to UpdateAUserNovelInBookShelfJsonHandler ")
	session, exist := c.GetQuery("session")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find session"}
		c.JSON(500, errJson)
		return
	}
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		c.JSON(500, errJson)
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return
	}
	//如果redis中有这个key的话那就再给他续一段时间
	//redis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))

	novelid, exist := c.GetQuery("novelid")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		return
	}
	nid, err := strconv.Atoi(novelid)
	if err != nil {
		errJson := JsonRet{Code: 0, Ret: "novelid err"}
		c.JSON(500, errJson)
		return
	}
	chapterName, exist := c.GetQuery("chaptername")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "can't find chapname"}
		c.JSON(500, errJson)
		return
	}
	chapterUrl, exist := c.GetQuery("chapterurl")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "can't find chapurl"}
		c.JSON(500, errJson)
		return
	}

	bookShelf := noveldb.BookShelf{}
	bookShelf.UserID = uid
	bookShelf.NovelID = nid
	bookShelf.ReadChapterName = chapterName
	bookShelf.ReadChapterUrl = chapterUrl
	noveldb.UpdateOneDataToBookShlefByUserIDAndNovelID(bookShelf)
	okJson := JsonRet{Code: 1, Ret: "update novel to bookshelf ok"}
	c.JSON(200, okJson)
}

func GetTheNovelInBookShelfJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTheNovelInBookShelfJsonHandler ")
	session, exist := c.GetQuery("session")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find session"}
		c.JSON(500, errJson)
		return
	}
	uid, err := myredis.GetRedisClient().Get(session).Result()
	if err == redis.Nil {
		errJson := JsonRet{Code: -1, Ret: "can't find uid, pls login"}
		c.JSON(500, errJson)
		return
	} else if err != nil {
		logger.ALogger().Errorf("Get Redis key Err :", err)
		panic(err)
		errJson := JsonRet{Code: -3, Ret: "panic"}
		c.JSON(500, errJson)
		return
	}
	//如果redis中有这个key的话那就再给他续一段时间
	//redis.REDIS_SAVE_TIME
	myredis.GetRedisClient().Expire(session, myredis.REDIS_SAVE_TIME)
	logger.ALogger().Debugf("Session : %s Refrash Time :%v", session, myredis.GetRedisClient().TTL(session))

	novelid, exist := c.GetQuery("novelid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		return
	}
	nid, err := strconv.Atoi(novelid)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "novelid err"}
		c.JSON(500, errJson)
		return
	}

	if _, find := noveldb.FindOneNovelFromBookShelfByUserIDAndNovelID(uid, nid); !find {
		okJson := JsonRet{Code: 0, Ret: "the novel not in your bookshelf"}
		c.JSON(200, okJson)
		return
	}

	okJson := JsonRet{Code: 1, Ret: "the novel in your bookshelf"}
	c.JSON(200, okJson)

	return
}
func AddANovelCommentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to AddANovelCommentJsonHandler ")
	return
}

func UpdateANovelCommentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to UpdateANovelCommentJsonHandler ")
	return
}
func GetANovelCommentsJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetANovelCommentsJsonHandler ")
	return
}
func DeleteANovelCommentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to DeleteANovelCommentJsonHandler ")
	return
}
