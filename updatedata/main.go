package main

import (
	"GoReadNovel/logger"
	"GoReadNovel/noveldb"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	MAX_PAGE = 336
)

func main() {
	logger.ALogger().Debug("Try to Update Data ")

	for page := 1; page <= MAX_PAGE; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTypeNovelList.py", "quanbu", "allvisit", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)
		//fmt.Println("输出:", str)

		datas := strings.Split(strings.TrimSpace(str), ",")

		for _, data := range datas {
			idUrlName := strings.Split(strings.TrimSpace(data), "--")
			//logger.ALogger().Debug("--------------", len(idUrlName))
			if len(idUrlName) != 9 {
				//logger.ALogger().Error("Get Python Error in Page ", strPage)
				//最后一个就是错的，所以要从这里跳过
				//if idUrlName != ""{
				//}
				continue
			}
			/*
				id, err := strconv.Atoi(idUrlName[0])
				if err != nil {
					//fmt.Println("这条数据有问题:", idUrlName[0], idUrlName[1])
					continue
				}
			*/
			novel := noveldb.Novel{}
			//novel.Index = id
			novel.NovelUrl = idUrlName[2] //"/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
			novel.NovelName = idUrlName[3]
			novel.LatestChpName = idUrlName[7]
			novel.Author = idUrlName[4]
			novel.Desc = idUrlName[5]
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			noveldb.InsertOneDataToNovel(novel)

		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 5s", page, MAX_PAGE)
		time.Sleep(5 * time.Second)
	}
}
