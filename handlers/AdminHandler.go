package handlers

import (
	_ "GoReadNovel/helpers"
	"GoReadNovel/logger"
	//"GoReadNovel/spider"
	"github.com/gin-gonic/gin"
	//"strings"
	"GoReadNovel/noveldb"
	// "fmt"
	"strconv"
)

func GetNovelTableInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNovelTableInfoJsonHandler")
	beginStr, exist := c.GetQuery("begin")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find begin"}
		c.JSON(500, errJson)
		logger.ALogger().Error("begin not find")
		return
	}
	numStr, exist := c.GetQuery("num")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find num"}
		c.JSON(500, errJson)
		logger.ALogger().Error("num not find")
		return
	}
	begin, _ := strconv.Atoi(beginStr)
	num, _ := strconv.Atoi(numStr)
	novelListMap, find := noveldb.FindDatasFromNovel(begin, num)
	if !find {
		errJson := JsonRet{Code: -2, Ret: "db find error"}
		c.JSON(500, errJson)
		logger.ALogger().Error("db find error")
		return
	}
	var novelsInfo []noveldb.Novel
	for i := 0; i < len(novelListMap); i++ {
		novelsInfo = append(novelsInfo, novelListMap[i])
	}
	okJson := JsonRet{Code: 1, Ret: novelsInfo}
	c.JSON(200, okJson)
	return
}

func GetEditNovelJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetEditNovelJsonHandler")
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	logger.ALogger().Debug("get from client:", string(buf[0:n]))
	//将json转换成struct 更新数据库

	return
}

func GetDeleteNovleIDHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetDeleteNovleIDHandler")
	novelid, exist := c.GetQuery("novelid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		logger.ALogger().Error("novelid not find")
		return
	}
	nid, err := strconv.Atoi(novelid)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "novelid is not a number"}
		c.JSON(500, errJson)
		return
	}
	//删除的代码先注释其他没有问题之后再投入使用
	/*
		if del := noveldb.DeleteOneDataToNovelByID(nid); !del {
			errJson := JsonRet{Code: 0, Ret: "db delete error"}
			c.JSON(500, errJson)
			logger.ALogger().Error("db delete error")
			return
		}
	*/
	okJson := JsonRet{Code: 1, Ret: "delete ok"}
	c.JSON(200, okJson)
	return
}

func GetNovelsCountHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNovelsCountHandler")

	count, get := noveldb.GetNovelsCountFromNovel()
	if !get {
		errJson := JsonRet{Code: 0, Ret: "cant get novel count."}
		c.JSON(500, errJson)
		return
	}
	okJson := JsonRet{Code: 1, Ret: count}
	c.JSON(200, okJson)
	return
}

func GetUltimateSearchNovelsJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetUltimateSearchNovelsJsonHandler")
	//按照什么搜索：ID、小说名称和作者、小说类型
	searchType, exist := c.GetQuery("type")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find type"}
		c.JSON(500, errJson)
		logger.ALogger().Error("type not find")
		return
	}
	keyStr, exist := c.GetQuery("key")
	if !exist {
		errJson := JsonRet{Code: -1, Ret: "can't find novelid"}
		c.JSON(500, errJson)
		logger.ALogger().Error("novelid not find")
		return
	}
	var novelListMap map[int]noveldb.Novel
	if searchType == "0" { //ID
		id, err := strconv.Atoi(keyStr)
		if err != nil {
			errJson := JsonRet{Code: -1, Ret: "key is not a number"}
			c.JSON(500, errJson)
			return
		}
		//获取到novel还是放到map中同一接口
		novelListMap = make(map[int]noveldb.Novel)
		novel, _ := noveldb.FindOneDataFromNovelByID(id)
		novelListMap[0] = novel

	} else if searchType == "1" { //name and author
		novelListmap, find = noveldb.FindDatasFromNovelByNameOrAuthor(key)
		if !find {
			//数据库中没有找到，之后可以添加使用爬虫爬取一下
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
	} else if searchType == "2" { //novel type
		novelListMap, find = noveldb.FindDatasFromNovelByNovelType(noveldb.NovelTypeEtoC[novelType])
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
	}
	//listMap的声明与处理
	okJson := JsonRet{Code: 1, Ret: novelListMap}
	c.JSON(200, okJson)
	return

}
