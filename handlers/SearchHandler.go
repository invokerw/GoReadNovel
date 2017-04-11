package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"GoReadNote/spider"
	"github.com/gin-gonic/gin"
	//"net/http"
	//"strings"
)

const (
	BYNOVELNAME = 0
	BYURL       = 1
)

func SearchNovelHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to SearchNovelHandler")
	h := gin.H{}
	novelName, exist := c.GetQuery("novelname")
	if !exist {
		c.JSON(500, h)
	}
	logger.ALogger().Notice(" novelname:", novelName)
	novelListMap, find := spider.SearchNovelByName(novelName)

	//logger.ALogger().Notice("Try to novelListMap:", novelListMap)
	if !find {
		//没有找到 再试试直接Get
		getNovelChapterList(c, novelName, "", BYNOVELNAME)
	} else {
		showSearchResult(c, novelListMap)
	}
	return
}
func GetBookInfoHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetBookInfoHandler")

	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
		return
	}
	name, exist := c.GetQuery("name")
	if !exist {
		c.JSON(500, h)
	}
	url = spider.URL + url
	logger.ALogger().Debug("url = ", url)

	getNovelChapterList(c, name, url, BYURL)
	return

}
func getNovelChapterList(c *gin.Context, name string, url string, getType int) {
	logger.ALogger().Debugf("Try to getNovelChapterList type = %d,name = %s,url = %s", getType, name, url)

	h := gin.H{}

	var chptMap map[int]spider.ChapterInfo
	var ok bool

	if getType == BYNOVELNAME {
		chptMap, ok = spider.GetNovelChapterListByNovelName(name)
	} else if getType == BYURL {
		chptMap, ok = spider.GetNovelChapterListByUrl(url)
	}

	if !ok {
		h["Title"] = "未知错误"
		helpers.Render(c, h, "err.tmpl")
		return
	}

	h["Title"] = "章节列表"
	h["Retname"] = name
	h["ChptList"] = chptMap

	helpers.Render(c, h, "novelchplist.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func showSearchResult(c *gin.Context, novelListMap map[int]spider.SearchNovel) {
	logger.ALogger().Debug("Try to showSearchResult")

	h := gin.H{}
	if len(novelListMap) == 0 {
		h["Title"] = "没有找到"
		h["Code"] = 0

	} else {
		h["Title"] = "搜索结果"
		h["Code"] = 1
	}
	h["NovelMap"] = novelListMap
	helpers.Render(c, h, "searchret.tmpl")
	return
}
