package handlers

import (
	"GoReadNovel/helpers"
	"GoReadNovel/logger"
	"GoReadNovel/spider"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func HomeHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to HomeHandler")
	//helpers.Render(c, gin.H{"Title": "首页"}, "index.tmpl")
	c.HTML(200, "index.html", gin.H{})
}
func NewHomeHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to NewHomeHandler")
	c.HTML(200, "index.html", gin.H{})
	//helpers.Render(c, gin.H{}, "index.html")
}
func GetSearchIndexHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetSearchIndexHandler")
	helpers.Render(c, gin.H{"Title": "搜索"}, "index.tmpl")
}
func GetNovelContentHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetNovelContentHandler")
	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)
	chp := spider.GetNovelContent(url)
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
	helpers.Render(c, h, "novel.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("%s\n\n%s\n", chp.ChapterName, chp.Content)))
	return
}

func WeiXinOnLoginHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to WeiXinOnLoginHandler")
	h := gin.H{}
	code, exist := c.GetQuery("code")
	userInfo, exist := c.GetQuery("info")
	if !exist {
		c.JSON(500, h)
		return
	}
	logger.ALogger().Debugf("code = %v,userInfo = %v", code, userInfo)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx9589545c06df6dab&secret=9fd41538a947f987781aebf457a2edc6&js_code="
	url = url + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)
	if err != nil {
		logger.ALogger().Error("Wx Server Get Error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logger.ALogger().Debug("body = ", body)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(body), &dat); err != nil {
		logger.ALogger().Error("Error = ", err)
	}
	logger.ALogger().Debug("Json = ", dat)
	//json =  map[session_key:H0nrxdNeQj674ze5kO+KAQ== expires_in:7200 openid:oRasZ0TOomboER5UC-KlkC_tGf20]
	//这里先只写返回正确的openid现象。事后还要加上不正确的时候
	type JsonHolder struct {
		OpenId     string `json:"opid"`
		SessionKey string `json:"sk"`
	}
	holder := JsonHolder{OpenId: dat["openid"].(string), SessionKey: dat["session_key"].(string)}
	c.JSON(200, holder)
	return
}
