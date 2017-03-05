package spider

import (
	"GoReadNote/logger"
	//"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type TopNote struct {
	Note
}

func GetTopNoteList() (map[int]TopNote, bool) {
	logger.ALogger().Debug("Try to GetTopNoteList ")

	cmd := exec.Command("python", "./spider/python/getTopNoteList.py")
	//cmd := exec.Command("python", "getTopNoteList.py")
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil, false
	}
	str := string(buf)
	//fmt.Println("输出:", str)
	var noteFindMap map[int]TopNote
	noteFindMap = make(map[int]TopNote)

	datas := strings.Split(strings.TrimSpace(str), ",")

	for _, data := range datas {
		idUrlName := strings.Split(strings.TrimSpace(data), "--")
		if len(idUrlName) != 7 {
			//fmt.Println("这个数据不为7:", idUrlName)
			continue
		}
		id, err := strconv.Atoi(idUrlName[0])
		if err != nil {
			//fmt.Println("这条数据有问题:", idUrlName[0], idUrlName[1])
			continue
		}

		sn := TopNote{}
		sn.Index = id
		sn.NoteUrl = "/GetBookInfo?go=" + idUrlName[1][len(URL):len(idUrlName[1])]
		sn.NoteName = idUrlName[2]
		sn.LatestChpName = idUrlName[3]
		sn.Author = idUrlName[4]
		sn.Status = idUrlName[5]
		sn.LatestChpUrl = "/GetBookInfo?go=" + idUrlName[6]

		noteFindMap[id] = sn

	}
	//fmt.Println("找到小说的数量:", len(noteFindMap))
	if len(noteFindMap) == 0 {
		return nil, false
	}
	return noteFindMap, true
}

/*
func main() {

	noteFindMap, _ := GetTopNoteList()
	for i := 1; i <= len(noteFindMap); i++ {
		fmt.Printf("%d : %v\n", i, noteFindMap[i])
	}

}
*/
