package spider

import (
	"GoReadNote/logger"
	//"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type TopTypeNote struct {
	Note
}

func GetTopByTypeNoteList(noteType string, sortType string, page string) (map[int]TopTypeNote, bool) {
	logger.ALogger().Debug("Try to GetTopTypeNoteList ")

	cmd := exec.Command("python", "./spider/python/getTopByTypeNoteList.py", noteType, sortType, page)

	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil, false
	}
	str := string(buf)
	//fmt.Println("输出:", str)
	var noteFindMap map[int]TopTypeNote
	noteFindMap = make(map[int]TopTypeNote)

	datas := strings.Split(strings.TrimSpace(str), ",")

	for _, data := range datas {
		idUrlName := strings.Split(strings.TrimSpace(data), "--")
		if len(idUrlName) != 8 {
			//fmt.Println("这个数据不为8:", idUrlName)
			continue
		}
		id, err := strconv.Atoi(idUrlName[0])
		if err != nil {
			//fmt.Println("这条数据有问题:", idUrlName[0], idUrlName[1])
			continue
		}

		sn := TopTypeNote{}
		sn.Index = id
		sn.NoteUrl = "/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
		sn.NoteName = idUrlName[3]
		sn.LatestChpName = idUrlName[7]
		sn.Author = idUrlName[4]
		sn.Desc = idUrlName[5]
		sn.LatestChpUrl = "/GetBookInfo?go=" + idUrlName[6]

		noteFindMap[id] = sn

	}
	logger.ALogger().Debug("找到小说的数量:", len(noteFindMap))
	if len(noteFindMap) == 0 {
		return nil, false
	}
	return noteFindMap, true
}

/*
func main() {

	noteFindMap, _ := GetTopByTypeNoteList()
	for i := 1; i <= len(noteFindMap); i++ {
		logger.ALogger().Debug("%d : %v\n", i, noteFindMap[i])
	}

}
*/
