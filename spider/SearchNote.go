package spider

import (
	"GoReadNote/logger"
	"os/exec"
	"strconv"
	"strings"
)

type SearchNote struct {
	Note
}

func SearchNoteByName(noteName string) (map[int]SearchNote, bool) {
	logger.ALogger().Debug("Try to SearchNoteByName noteName:", noteName)

	cmd := exec.Command("python", "./spider/python/searchNote.py", noteName)
	//cmd := exec.Command("python", "searchNote.py", noteName)
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil, false
	}
	str := string(buf)
	//fmt.Println("输出:", str)
	var noteFindMap map[int]SearchNote
	noteFindMap = make(map[int]SearchNote)

	datas := strings.Split(strings.TrimSpace(str), ",")

	for _, data := range datas {
		idUrlName := strings.Split(strings.TrimSpace(data), "--")
		if len(idUrlName) != 7 {
			//fmt.Println("这个数据不为3:", idUrlName)
			continue
		}
		id, err := strconv.Atoi(idUrlName[0])
		if err != nil {
			//fmt.Println("这条数据有问题:", idUrlName[0], idUrlName[1])
			continue
		}

		sn := SearchNote{}
		sn.Index = id
		sn.NoteUrl = "/GetBookInfo?go=" + idUrlName[1][len(URL):len(idUrlName[1])]
		sn.NoteName = idUrlName[2]
		sn.LatestChpName = idUrlName[3]
		sn.Author = idUrlName[4]
		sn.Status = idUrlName[5]
		sn.LatestChpUrl = "/GetBookInfo?go=" + idUrlName[6]
		noteFindMap[id] = sn

	}
	//logger.ALogger().Debug("找到小说的数量:", len(noteFindMap))
	if len(noteFindMap) == 0 {
		return nil, false
	}
	return noteFindMap, true
}

/*
func main() {

	noteFindMap, _ := SearchNoteByName("遮天")
	for i := 1; i <= len(noteFindMap); i++ {
		fmt.Printf("%d : %v\n", i, noteFindMap[i])
	}

}
*/