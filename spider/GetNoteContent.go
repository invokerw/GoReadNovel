package spider

import (
	"GoReadNote/logger"
	"os/exec"
	"strings"
)

type ChapterContent struct {
	Content     string `json:"content"`  //章节内容
	ChapterName string `json:"chpname"`  //章节名称
	Url         string `json:"churl"`    //地址
	ChpNext     string `json:"nexturl"`  //上一章
	ChpPre      string `json:"preurl"`   //下一章
	NoteName    string `json:"notename"` //小说名字
}

func GetNoteContent(url string) *ChapterContent {
	logger.ALogger().Debug("Try to GetNoteContent url:", url)

	cmd := exec.Command("python", "./spider/python/getNoteContent.py", url)
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil
	}
	str := string(buf)

	datas := strings.Split(strings.TrimSpace(str), "-|-")
	if len(datas) != 5 {
		//fmt.Println("这个数据不为2:", datas)
		return nil
	}
	//logger.ALogger().Debug("datas:", datas)
	cc := ChapterContent{}
	cc.Content = datas[1]
	cc.Url = url
	cc.NoteName = datas[4]
	cc.ChapterName = datas[0]
	if datas[2][len(datas[2])-1] == '/' {
		cc.ChpPre = "/GetBookInfo?go=" + datas[2] + "&name=" + cc.NoteName
	} else {
		cc.ChpPre = "/GetBookContent?go=" + datas[2]
	}
	cc.ChpNext = "/GetBookContent?go=" + datas[3]
	//fmt.Println(idUrlName[2])

	return &cc
}

/*
func main() {
	url := "http://www.huanyue123.com/book/0/11/2925296.html"
	chc := GetNoteContent(url)
	fmt.Println(chc.ChapterName)
	fmt.Println(chc.Content)

}
*/
