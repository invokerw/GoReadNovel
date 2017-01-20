package handlers

import (
	"GoReadNote/helpers"
	"GoReadNote/logger"
	"GoReadNote/sprider"
	"github.com/gin-gonic/gin"
	//"net/http"
	//"strings"
)

const (
	BYNOTENAME = 0
	BYURL      = 1
)

func SearchNoteHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to SearchNoteHandler")
	h := gin.H{}
	noteName, exist := c.GetQuery("notename")
	if !exist {
		c.JSON(500, h)
	}
	noteListMap, find := sprider.SearchNoteByName(noteName)
	if !find {
		//没有找到 再试试直接Get
		getNoteChapterList(c, noteName, "", BYNOTENAME)
	} else {
		showSearchResult(c, noteListMap)
	}
	return
}
func GetBookInfoHandler(c *gin.Context) {
	logger.ALogger().Debug("Try to GetBookInfoHandler")

	h := gin.H{}
	url, exist := c.GetQuery("go")
	if !exist {
		c.JSON(500, h)
	}
	name, exist := c.GetQuery("name")
	if !exist {
		c.JSON(500, h)
	}
	url = sprider.URL + url
	logger.ALogger().Debug("url = ", url)

	getNoteChapterList(c, name, url, BYURL)

}
func getNoteChapterList(c *gin.Context, name string, url string, getType int) {
	logger.ALogger().Debugf("Try to getNoteChapterList type = %d,name = %s,url = %s", getType, name, url)

	h := gin.H{}

	var chptMap map[int]sprider.ChapterInfo
	var ok bool

	if getType == BYNOTENAME {
		chptMap, ok = sprider.GetNoteChapterListByNoteName(name)
	} else if getType == BYURL {
		chptMap, ok = sprider.GetNoteChapterListByUrl(url)
	}

	if !ok {
		h["Title"] = "未知错误"
		helpers.Render(c, h, "err.tmpl")
		return
	}

	h["Title"] = "章节列表"
	h["Retname"] = name
	h["ChptList"] = chptMap

	helpers.Render(c, h, "notechplist.tmpl")
	//c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func showSearchResult(c *gin.Context, noteListMap map[int]sprider.SearchNote) {
	logger.ALogger().Debug("Try to showSearchResult")

	h := gin.H{}
	if len(noteListMap) == 0 {
		h["Title"] = "没有找到"
		h["Code"] = 0

	} else {
		h["Title"] = "搜索结果"
		h["Code"] = 1
	}
	h["NoteMap"] = noteListMap
	helpers.Render(c, h, "searchret.tmpl")
	return
}
