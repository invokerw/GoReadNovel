package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeHandler(c *gin.Context) {
	logger.ALogger().Debug("Index.")
	helpers.Render(c, gin.H{"Title": "首页"}, "index.tmpl")
}
func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
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
