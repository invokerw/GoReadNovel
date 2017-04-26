package spider

import (
	"GoReadNovel/logger"
	_ "fmt"
	"os/exec"
	"strconv"
	"strings"
)

type ChapterInfo struct {
	Index       int    `json:"index"`  //第几章索引
	ChapterName string `json:"chname"` //章节名称
	Url         string `json:"churl"`  //地址
}

/*
	//是否加携程会好一些
func GoroutineGetNovelChapterListByNovelName(ch chan map[int]ChapterInfo, novelName string) {
	ch <- GetNovelChapterListByNovelName(novelName)
}
*/

func GetNovelChapterListByNovelName(novelName string) (map[int]ChapterInfo, bool) {
	logger.ALogger().Debug("Try to GetNovelChapterListByNovelName NovelName:", novelName)
	cmd := exec.Command("python", "./python/getNovelChaptersBySearch.py", novelName)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println("%v", err)
		logger.ALogger().Error(err)
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

func GetNovelChapterListByUrl(url string) (map[int]ChapterInfo, bool) {
	logger.ALogger().Debug("Try to GetNovelChapterListByUrl url:", url)
	cmd := exec.Command("python", "./python/getNovelChaptersByUrl.py", url)
	buf, err := cmd.Output()
	if err != nil {
		//fmt.Println("%v", err)
		logger.ALogger().Error(err)
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
		//应该把"/GetBookContent?go="删掉
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

	chpMap = GetNovelChapterList("圣墟")
	for i := 1; i <= len(chpMap); i++ {
		fmt.Printf("第%d章:%s\n", i, chpMap[i].ChapterName)
	}

}
*/
