package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	//"GoReadNote/sprider"
	"github.com/gin-gonic/gin"
	"net/http"
	//"strings"
	"path/filepath"
	"os"
	"fmt"
	"io"
)
const(
	Upload_Dir = "./savefile/"
)
func GetUpLoadPageHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to HomeHandler")
	helpers.Render(c, gin.H{"Title": "文件上传"}, "uploadfile.tmpl")
}

func UploadFileHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to UploadFileHandler")
	//h := gin.H{}

	//在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
	//Max 16MB
	c.Request.ParseMultipartForm(2 << 24) 

	file, handler, err := c.Request.FormFile("uploadfile")

	if err != nil{
		//上传错误
		fmt.Printf("上传错误")
		return
	}
	//check file type

	fileName := handler.Filename
	f, _ := os.OpenFile(Upload_Dir+fileName, os.O_CREATE|os.O_WRONLY, 0660)
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Printf("上传失败")
		return
	}
	filedir, _ := filepath.Abs(Upload_Dir + fileName)
	fmt.Printf(fileName+"上传完成,服务器地址:"+filedir)	
	//helpers.Render(c, h, "note.tmpl")
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("%s\n","OK")))
	return
}
