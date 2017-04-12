package spider

import (
	"GoReadNovel/logger"
	//"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type TopNovel struct {
	Novel
}

func GetTopNovelList() (map[int]TopNovel, bool) {
	logger.ALogger().Debug("Try to GetTopNovelList ")

	cmd := exec.Command("python", "./python/getTopNovelList.py")
	//cmd := exec.Command("python", "getTopNovelList.py")
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Error("%v", err)
		return nil, false
	}
	str := string(buf)
	//fmt.Println("输出:", str)
	var novelFindMap map[int]TopNovel
	novelFindMap = make(map[int]TopNovel)

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

		sn := TopNovel{}
		sn.Index = id
		sn.NovelUrl = "/GetBookInfo?go=" + idUrlName[1][len(URL):len(idUrlName[1])]
		sn.NovelName = idUrlName[2]
		sn.LatestChpName = idUrlName[3]
		sn.Author = idUrlName[4]
		sn.Status = idUrlName[5]
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

	novelFindMap, _ := GetTopNovelList()
	for i := 1; i <= len(novelFindMap); i++ {
		fmt.Printf("%d : %v\n", i, novelFindMap[i])
	}

}
*/
