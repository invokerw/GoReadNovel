package handlers

import (
	"GoReadNovel/config"
	_ "GoReadNovel/helpers"
	"GoReadNovel/logger"
	//"GoReadNovel/spider"
	"GoReadNovel/noveldb"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os/exec"
	"strconv"
	"strings"
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
	pagecount, err := config.GetPythonConfigInterface().Int("pagecount", "pagecount")
	logger.ALogger().Debugf("pagecount = %d", pagecount)
	//pagecount, err := strconv.Atoi(pagecountStr)
	if err != nil {
		logger.ALogger().Error("GetSpiderConfigJsonHandler get pagecount config error:", err)
		//errJson := JsonRet{Code: -1, Ret: "get pagecount error"}
		//c.JSON(500, errJson)
		//return
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
	confJson, exist := c.GetQuery("config")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find confJson"}
		c.JSON(500, errJson)
		return
	}
	//将json转换成struct
	var conf config.PythonConfig
	json.Unmarshal([]byte(confJson), &conf)
	//logger.ALogger().Debug("get from client confJson:", conf)

	cmd := exec.Command("rm", "./python/test.conf")
	_, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("rm err = ", err)
	}
	cmd = exec.Command("cp", "./python/python.conf", "./python/test.conf")
	_, err = cmd.Output()
	if err != nil {
		logger.ALogger().Error("cp err = ", err)
	}
	config.TestConfigReload()
	if !config.WriteATestPageConfig(conf) {
		errJson := JsonRet{Code: -1, Ret: "WriteATestPageConfig err"}
		c.JSON(500, errJson)
		return
	}
	//这里还需要改
	cmdStr := "python ./python/test_" + conf.Name + " test.conf"
	for i := 0; i < conf.TestValueCount; i++ {
		cmdStr = cmdStr + " " + conf.TestValues[i]
	}
	logger.ALogger().Debug("Exec CMD :", cmdStr)
	list := strings.Split(cmdStr, " ")
	//cmd = exec.Command("python", "./python/test_getMaxPageNum.py", "test.conf")
	cmd = exec.Command(list[0], list[1:]...)
	//getTopByTypeNovelList.py quanbu allvisit 1")
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("py err = ", err)
	}
	str := string(buf)
	logger.ALogger().Debug("python ret str = ", str)
	okJson := JsonRet{Code: 1, Ret: str}
	c.JSON(200, okJson)
	return
}

//先经过上面的测试通过之后才能保存。
func SaveConfigJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to SaveConfigJsonHandler")
	cmd := exec.Command("mv", "./python/python.conf", "./python/old_python.conf")
	_, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("mv err = ", err)
		errJson := JsonRet{Code: -2, Ret: "mv err"}
		c.JSON(500, errJson)
		return

	}
	cmd = exec.Command("cp", "./python/test.conf", "./python/python.conf")
	_, err = cmd.Output()
	if err != nil {
		logger.ALogger().Error("cp err = ", err)
		errJson := JsonRet{Code: -1, Ret: "cp err"}
		c.JSON(500, errJson)
		return
	}
	okJson := JsonRet{Code: 1, Ret: "ok"}
	c.JSON(200, okJson)
	return
}

func GetUserFeedbackJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetUserFeedbackJsonHandler")
	feedbacktypeStr, exist := c.GetQuery("feedbacktype")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find feedbacktype"}
		c.JSON(500, errJson)
		logger.ALogger().Error("feedbacktype not find")
		return
	}
	solvedStr, exist := c.GetQuery("solved")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find solved"}
		c.JSON(500, errJson)
		logger.ALogger().Error("solved not solved")
		return
	}
	feedbacktype, err := strconv.Atoi(feedbacktypeStr)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "feedbacktype atoi fail"}
		c.JSON(500, errJson)
		logger.ALogger().Error("feedbacktype atoi fail")
		return
	}
	solved, err := strconv.Atoi(solvedStr)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "solved atoi fail"}
		c.JSON(500, errJson)
		logger.ALogger().Error("solved atoi fail")
		return
	}
	feedbacks, find := noveldb.FindDatasFromFeedback(feedbacktype, solved)
	if !find {
		errJson := JsonRet{Code: 0, Ret: "not find feedback"}
		c.JSON(500, errJson)
		return
	}
	var fbs []noveldb.Feedback
	for i := 0; i < len(feedbacks); i++ {
		feedback := noveldb.Feedback{}
		feedback = feedbacks[i]
		fbs = append(fbs, feedback)
	}
	okJson := JsonRet{Code: 1, Ret: fbs}
	c.JSON(200, okJson)
	return
}

func SolvedUserFeedbackJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to SolvedUserFeedbackJsonHandler")
	contactidStr, exist := c.GetQuery("contactid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find contactid"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid not find")
		return
	}

	contactid, err := strconv.Atoi(contactidStr)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "contactid atoi fail"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid atoi fail")
		return
	}
	if _, find := noveldb.FindOneDataFromFeedbackByFeedbackID(contactid); !find {
		errJson := JsonRet{Code: 0, Ret: "contactid not faind in sql"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid not faind in sql")
		return
	}
	noveldb.UpdateOneDataSolvedToFeedbackByFeedbackID(contactid)
	okJson := JsonRet{Code: 1, Ret: "ok"}
	c.JSON(200, okJson)
}

func DelAUserFeedbackJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to DelAUserFeedbackJsonHandler")
	contactidStr, exist := c.GetQuery("contactid")
	if !exist {
		errJson := JsonRet{Code: -2, Ret: "can't find contactid"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid not find")
		return
	}

	contactid, err := strconv.Atoi(contactidStr)
	if err != nil {
		errJson := JsonRet{Code: -1, Ret: "contactid atoi fail"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid atoi fail")
		return
	}
	if _, find := noveldb.FindOneDataFromFeedbackByFeedbackID(contactid); !find {
		errJson := JsonRet{Code: 0, Ret: "contactid not faind in sql"}
		c.JSON(500, errJson)
		logger.ALogger().Error("contactid not faind in sql")
		return
	}
	noveldb.DeleteOneDataToFeedbackByID(contactid)
	okJson := JsonRet{Code: 1, Ret: "ok"}
	c.JSON(200, okJson)
}
