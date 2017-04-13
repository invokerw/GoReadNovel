package spider

import (
	"GoReadNovel/logger"
	//"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type TopTypeNovel struct {
	Novel
}

func GetTopByTypeNovelList(novelType string, sortType string, page string) (map[int]TopTypeNovel, bool) {
	logger.ALogger().Debug("Try to GetTopTypeNovelList ")

	cmd := exec.Command("python", "./python/getTopByTypeNovelList.py", novelType, sortType, page)

	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil, false
	}
	str := string(buf)
	//fmt.Println("输出:", str)
	var novelFindMap map[int]TopTypeNovel
	novelFindMap = make(map[int]TopTypeNovel)

	datas := strings.Split(strings.TrimSpace(str), ",")

	for _, data := range datas {
		idUrlName := strings.Split(strings.TrimSpace(data), "--")
		if len(idUrlName) != 9 {
			//fmt.Println("这个数据不为9:", idUrlName)
			continue
		}
		id, err := strconv.Atoi(idUrlName[0])
		if err != nil {
			//fmt.Println("这条数据有问题:", idUrlName[0], idUrlName[1])
			continue
		}

		sn := TopTypeNovel{}
		sn.Index = id
		sn.NovelUrl = "/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
		sn.NovelName = idUrlName[3]
		sn.LatestChpName = idUrlName[7]
		sn.Author = idUrlName[4]
		sn.Desc = idUrlName[5]
		sn.LatestChpUrl = "/GetBookInfo?go=" + idUrlName[6]

		novelFindMap[id] = sn

	}
	logger.ALogger().Debug("找到小说的数量:", len(novelFindMap))
	if len(novelFindMap) == 0 {
		return nil, false
	}
	return novelFindMap, true
}

/*
func main() {

	novelFindMap, _ := GetTopByTypeNovelList()
	for i := 1; i <= len(novelFindMap); i++ {
		logger.ALogger().Debug("%d : %v\n", i, novelFindMap[i])
	}

}
*/
