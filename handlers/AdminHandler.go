package handlers

import (
	"GoReadNovel/config"
	_ "GoReadNovel/helpers"
	"GoReadNovel/logger"
	//"GoReadNovel/spider"
	"github.com/gin-gonic/gin"
	//"strings"
	"GoReadNovel/noveldb"
	"encoding/json"
	"os/exec"
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
	novelJson, exist := c.GetQuery("novel")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novel"}
		c.JSON(500, errJson)
		return
	}
	//将json转换成struct 更新数据库
	var novel noveldb.Novel
	//logger.ALogger().Debugf("novelJson:%v", novelJson)
	json.Unmarshal([]byte(novelJson), &novel)
	logger.ALogger().Debug("get from client novel:", novel)
	noveldb.UpdateOneDataToNovelByID(novel)
	okJson := JsonRet{Code: 1, Ret: "ok"}
	c.JSON(200, okJson)
	return
}
func GetANewNovelJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetANewNovelJsonHandler")
	novelJson, exist := c.GetQuery("novel")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find novel"}
		c.JSON(500, errJson)
		return
	}
	//将json转换成struct 更新数据库
	var novel noveldb.Novel
	json.Unmarshal([]byte(novelJson), &novel)
	logger.ALogger().Debug("get from client novel:", novel)
	noveldb.InsertOneDataToNovel(novel)
	okJson := JsonRet{Code: 1, Ret: "ok"}
	c.JSON(200, okJson)
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
	logger.ALogger().Debug("delete novel id : ", nid)
	//删除的代码先注释其他没有问题之后再投入使用

	if del := noveldb.DeleteOneDataToNovelByID(nid); !del {
		errJson := JsonRet{Code: 0, Ret: "db delete error"}
		c.JSON(500, errJson)
		logger.ALogger().Error("db delete error")
		return
	}

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
	var find bool
	if searchType == "0" {
		id, err := strconv.Atoi(keyStr)
		if err != nil { //name and author
			novelListMap, find = noveldb.FindDatasFromNovelByNameOrAuthor(keyStr)
			if !find {
				errJson := JsonRet{Code: 0, Ret: "can't find"}
				c.JSON(500, errJson)
				return
			}
		} else { //ID
			//获取到novel还是放到map中同一接口
			novelListMap = make(map[int]noveldb.Novel)
			novel, _ := noveldb.FindOneDataFromNovelByID(id)
			novelListMap[0] = novel
		}

	} else if searchType == "1" { //name and author
		/*	novelListMap, find = noveldb.FindDatasFromNovelByNameOrAuthor(keyStr)
			if !find {
				errJson := JsonRet{Code: 0, Ret: "can't find"}
				c.JSON(500, errJson)
				return
			}
		} else if searchType == "2" {*/ //novel type
		//限制数量或者不限制
		//novelListMap, find = noveldb.FindDatasFromNovelByNovelTypeNoLimitCount(keyStr)
		novelListMap, find = noveldb.FindDatasFromNovelByNovelType(keyStr)
		if !find {
			errJson := JsonRet{Code: 0, Ret: "can't find"}
			c.JSON(500, errJson)
			return
		}
	}
	//listMap的声明与处理
	var novelsInfo []noveldb.Novel
	for i := 0; i < len(novelListMap); i++ {
		novelsInfo = append(novelsInfo, novelListMap[i])
	}
	okJson := JsonRet{Code: 1, Ret: novelsInfo}
	c.JSON(200, okJson)
	return

}

func GetUsersInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetUsersInfoJsonHandler")
	usersMap, find := noveldb.FindDatasFromUser(0, 100)
	if !find {
		errJson := JsonRet{Code: 0, Ret: "can't find"}
		c.JSON(500, errJson)
		return
	}
	var usersInfo []noveldb.User
	for i := 0; i < len(usersMap); i++ {
		usersInfo = append(usersInfo, usersMap[i])
	}
	okJson := JsonRet{Code: 1, Ret: usersInfo}
	c.JSON(200, okJson)
	return
}
func GetSpiderConfigJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSpiderConfigJsonHandler")
	//str1 := config.GetPythonConfigInterface().String("getmaxpagenum::text1")
	//logger.ALogger().Debug("getmaxpagenum::text1 = ", str1)
	pagecount, err := config.GetPythonConfigInterface().Int("count::pagecount")
	if err != nil {
		logger.ALogger().Error("GetSpiderConfigJsonHandler get pagecount config error:", err)
		errJson := JsonRet{Code: -1, Ret: "get pagecount error"}
		c.JSON(500, errJson)
		return
	}
	var confs []config.PythonConfig
	for i := 0; i < pagecount; i++ {
		conf := config.PythonConfig{}
		conf, get := config.GetAPythonPageConfig(i)
		if !get {
			errJson := JsonRet{Code: 0, Ret: "get config error"}
			c.JSON(500, errJson)
			return
		}
		confs = append(confs, conf)
	}
	okJson := JsonRet{Code: 1, Ret: confs}
	c.JSON(200, okJson)
	return
}
func TestConfigJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to TestConfigJsonHandler")

	cmd := exec.Command("python", "./python/test_getMaxPageNum.py", "python.conf")
	//getTopByTypeNovelList.py quanbu allvisit 1")
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("err = ", err)
	}
	str := string(buf)
	logger.ALogger().Debug("str = ", str)
	return
}
func SaveConfigJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to SaveConfigJsonHandler")

	return
}
