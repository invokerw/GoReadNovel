package sprider

import (
	"GoReadNote/logger"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type ChapterInfo struct {
	Index       int    //第几章索引
	ChapterName string //章节名称
	Url         string //地址
}

const (
	URL = "http://www.huanyue123.com"
)

func GetNoteChapterList(noteName string) (map[int]ChapterInfo, bool) {
	logger.ALogger().Debug("Try to GetNoteChapterList noteName:", noteName)
	cmd := exec.Command("python", "./sprider/getNoteChapters.py", noteName)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("%v", err)
		return nil, false
	}
	str := string(buf)
	var chpMap map[int]ChapterInfo
	chpMap = make(map[int]ChapterInfo)

	datas := strings.Split(strings.TrimSpace(str), ",")
	for _, data := range datas {
		idUrlName := strings.Split(strings.TrimSpace(data), "-")
		if len(idUrlName) != 3 {
			//fmt.Println("这个数据不为3:", idUrlName)
			continue
		}
		id, err := strconv.Atoi(idUrlName[0])
		if err != nil {
			//fmt.Println("这个Url有问题:", idUrlName[0], idUrlName[1])
			continue
		}

		cp := ChapterInfo{}
		cp.Index = id
		cp.Url = "/GetBookContent?go=" + idUrlName[1][len(URL):len(idUrlName[1])]
		cp.ChapterName = idUrlName[2]
		//fmt.Println(idUrlName[2])
		chpMap[id] = cp

	}
	//fmt.Println("小说的章数:", len(chpMap))
	return chpMap, true
}

/*
func main() {

	chpMap = GetNoteChapterList("圣墟")
	for i := 1; i <= len(chpMap); i++ {
		fmt.Printf("第%d章:%s\n", i, chpMap[i].ChapterName)
	}

}
*/