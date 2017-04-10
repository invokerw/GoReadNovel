package handlers

import (
	"GoReadNote/logger"
	"GoReadNote/spider"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strconv"
)

//返回Json的一个模板  Code在不同情况下有不同作用
type JsonRet struct {
	Code int         `json:"code"`
	List interface{} `json:"list"`
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

func GetSearchNoteJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNoteJsonHandler")

	h := gin.H{}

	noteName, exist := c.GetQuery("notename")
	if !exist {
		c.JSON(500, h)
		return
	}
	noteListMap, find := spider.SearchNoteByName(noteName)
	if !find {
		//没有找到 再试试直接Get
		chptMap, ok := spider.GetNoteChapterListByNoteName(noteName)
		if !ok {
			c.JSON(500, h)
			return
		}
		var cpInfo []spider.ChapterInfo

		for i := 1; i <= len(chptMap); i++ {
			cpInfo = append(cpInfo, chptMap[i])
		}
		//code = 0 为一个结果  code = 1为小说列表
		retJson := JsonRet{Code: 0, List: cpInfo}
		c.JSON(200, retJson)
		return
	}
	var noteInfo []spider.SearchNote

	for i := 1; i <= len(noteListMap); i++ {
		noteInfo = append(noteInfo, noteListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return
}

func GetBookContentJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchNoteJsonHandler")
	h := gin.H{}

	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chp := spider.GetNoteContent(url)
	if chp == nil {
		c.JSON(500, h)
		return
	}

	c.JSON(200, chp)
	return

}

func GetTopNoteListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopNoteListJsonHandler")
	h := gin.H{}
	type JsonRet struct {
		Code int         `json:"code"` //code = 0 为一个结果  code = 1为小说列表
		List interface{} `json:"list"`
	}
	noteListMap, ok := spider.GetTopNoteList()
	if !ok {
		c.JSON(500, h)
		return
	}
	var noteInfo []spider.TopNote

	for i := 1; i <= len(noteListMap); i++ {
		noteInfo = append(noteInfo, noteListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return

}

func GetTopByTypeNoteListJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetTopByTypeNoteListJsonHandler")
	h := gin.H{}
	type JsonRet struct {
		Code int         `json:"code"` //code = 0 为一个结果  code = 1为小说列表
		List interface{} `json:"list"`
	}
	//获取
	noteType, exist := c.GetQuery("ntype")
	if !exist {
		noteType = "quanbu"
	}
	sortType, exist := c.GetQuery("stype")
	if !exist {
		sortType = "default"
	}
	page, exist := c.GetQuery("page")
	if !exist {
		page = "1"
	}

	noteListMap, ok := spider.GetTopByTypeNoteList(noteType, sortType, page)
	if !ok {
		c.JSON(500, h)
		return
	}
	var noteInfo []spider.TopTypeNote

	for i := 1; i <= len(noteListMap); i++ {
		noteInfo = append(noteInfo, noteListMap[i])
	}
	//code = 0 为一个结果  code = 1为小说列表
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return
}

func GetNoteInfoJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNoteInfoJsonHandler")
	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chptMap, ok := spider.GetNoteChapterListByUrl(url)
	if !ok {
		c.JSON(500, h)
		return
	}
	var noteInfo []spider.ChapterInfo

	for i := 1; i <= len(chptMap); i++ {
		noteInfo = append(noteInfo, chptMap[i])
	}
	retJson := JsonRet{Code: 1, List: noteInfo}
	c.JSON(200, retJson)
	return

}

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
