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

func GetTNovelTableInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTNovelTableInfoJsonHandler")
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
