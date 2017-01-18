package sprider

import (
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

func GetNoteChapterList(noteName string) map[int]ChapterInfo {
	cmd := exec.Command("python", "getNoteChapters.py", noteName)
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println("%v", err)
		return nil
	}
	str := string(buf)
	var chpMap map[int]Chapter
	chpMap = make(map[int]Chapter)

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
		cp.Url = idUrlName[1]
		cp.ChapterName = idUrlName[2]
		//fmt.Println(idUrlName[2])
		chpMap[id] = cp

	}
	//fmt.Println("小说的章数:", len(chpMap))
	return chpMap
}

/*
func main() {

	chpMap = GetNoteChapterList("圣墟")
	for i := 1; i <= len(chpMap); i++ {
		fmt.Printf("第%d章:%s\n", i, chpMap[i].ChapterName)
	}

}
*/
