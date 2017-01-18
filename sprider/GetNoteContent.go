package sprider

import (
	"fmt"
	"os/exec"
	"strings"
)

type ChapterContent struct {
	Content     string //章节内容
	ChapterName string //章节名称
	Url         string //地址
}

func GetNoteContent(url string) *ChapterContent {
	cmd := exec.Command("python", "getNoteContent.py", url)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("%v", err)
		return nil
	}
	str := string(buf)

	datas := strings.Split(strings.TrimSpace(str), "-|-")
	if len(datas) != 2 {
		//fmt.Println("这个数据不为2:", datas)
		return nil
	}

	cc := ChapterContent{}
	cc.Content = datas[1]
	cc.Url = url
	cc.ChapterName = datas[0]
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
