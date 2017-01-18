package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
	h["Title"] = "搜索结果"
	h["retname"] = noteName

	logger.ALogger().Debugf("noteName : %s", noteName)
	helpers.Render(c, h, "searchret.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
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
