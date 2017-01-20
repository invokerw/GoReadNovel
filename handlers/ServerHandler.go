package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"GoReadNote/sprider"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HomeHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to HomeHandler")
	helpers.Render(c, gin.H{"Title": "首页"}, "index.tmpl")
}
func GetNoteChapterListByNoteNameHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNoteChapterListByNoteNameHandler")
	h := gin.H{}

	noteName, exist := c.GetQuery("notename")
	if !exist {
		c.JSON(500, h)
	}
	chptMap, ok := sprider.GetNoteChapterList(noteName)

	if !ok {
		h["Title"] = "未知错误"
		helpers.Render(c, h, "err.tmpl")
		return
	}

	h["Title"] = "搜索结果"
	h["Retname"] = noteName
	h["ChptList"] = chptMap

	logger.ALogger().Debugf("noteName : %s", noteName)
	helpers.Render(c, h, "searchret.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func GetNoteContentHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNoteContentHandler")
	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
	}
	url = sprider.URL + url
	logger.ALogger().Debug("url = ", url)
	chp := sprider.GetNoteContent(url)
	if chp == nil {
		h["Title"] = "未知错误"
		helpers.Render(c, h, "err.tmpl")
		return
	}
	h["Title"] = chp.ChapterName
	//chp.Content = strings.Replace(chp.Content, "\n", "<br/>", -1) //字符串替换 指定起始位置为小于0,则全部替换 f00
	h["Chapter"] = chp
	h["ContentArry"] = strings.Split(strings.TrimSpace(chp.Content), "\n")
	helpers.Render(c, h, "note.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("%s\n\n%s\n", chp.ChapterName, chp.Content)))
	return
}
func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "123"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}
