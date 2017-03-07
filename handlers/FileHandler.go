package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	//"GoReadNote/spider"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strings"
	_ "fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	Upload_Dir = "./savefile/"
)

func GetUpLoadPageHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to HomeHandler")
	helpers.Render(c, gin.H{"Title": "文件上传"}, "uploadfile.tmpl")
}

func UploadFileHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to UploadFileHandler")
	h := gin.H{}
	//ftype = 0 ./savafile  1
	ftype, exist := c.GetPostForm("ftype")
	logger.ALogger().Debugf(ftype)
	if !exist {
		c.JSON(500, h)
		logger.ALogger().Error("没有发现ftype")
		return
	}
	//在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
	//Max 16MB
	c.Request.ParseMultipartForm(2 << 24)

	file, handler, err := c.Request.FormFile("uploadfile")

	if err != nil {
		//上传错误
		logger.ALogger().Error("上传错误")
		c.JSON(500, h)
		return
	}
	//check file type

	fileName := handler.Filename

	if ftype == "main" {

		filedir, _ := filepath.Abs(Upload_Dir + "main/" + fileName)
		f, _ := os.OpenFile(filedir, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			logger.ALogger().Error("上传失败")
			c.JSON(500, h)
			return
		}
		logger.ALogger().Debug(fileName + "上传完成,服务器地址:" + filedir)
		c.Redirect(http.StatusMovedPermanently, "/GetFileList")
	} else if ftype == "wei" {

		filedir, _ := filepath.Abs(Upload_Dir + "wei/" + fileName)
		f, _ := os.OpenFile(filedir, os.O_CREATE|os.O_WRONLY, 0660)
		_, err = io.Copy(f, file)
		if err != nil {
			logger.ALogger().Error("上传失败")
			c.JSON(500, h)
			return
		}
		logger.ALogger().Debug(fileName + "上传完成,服务器地址:" + filedir)
		c.Redirect(http.StatusMovedPermanently, "/Weifei")
	}
	return
}
