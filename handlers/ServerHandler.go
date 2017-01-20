package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"GoReadNote/sprider"
	"github.com/gin-gonic/gin"
	//"net/http"
	"strings"
)

func HomeHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to HomeHandler")
	helpers.Render(c, gin.H{"Title": "首页"}, "index.tmpl")
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
	//chp.Content = strings.Replace(chp.Content, "\n", "<br/>", -1) //字符串替换 指定起始位置为小于0,则全部替换
	h["Chapter"] = chp
	//logger.ALogger().Notice("chp.Content:", chp.Content)

	h["ContentArry"] = strings.Split(strings.TrimSpace(chp.Content), "\n")
	helpers.Render(c, h, "note.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("%s\n\n%s\n", chp.ChapterName, chp.Content)))
	return
}

func GetJsonHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetJsonHandler")
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	holder := JsonHolder{Id: 1, Name: "123"}
	//若返回json数据，可以直接使用gin封装好的JSON方法

	logger.ALogger().Debug("This")
	c.JSON(200, holder)

	return
}
