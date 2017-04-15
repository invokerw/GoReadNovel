package main

import (
	"GoReadNovel/logger"
	"GoReadNovel/noveldb"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	MAX_PAGE = 336
)

func InsertData() {
	logger.ALogger().Debug("Try to Insert Data ")

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
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)
	}
}
func UpdateData() {
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
			//存在novel就更新数据 不存在就插入一条新数据
			if _, exit := noveldb.FindOneDataByNovelNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
			} else {
				cmd = exec.Command("python", "../python/getNovelInfo.py", novel.NovelUrl)
				buf, err := cmd.Output()
				if err != nil {
					logger.ALogger().Errorf("Novel %s Get Info Error 1,Url %s", novel.NovelName, novel.NovelUrl)
					continue
				}

				str = string(buf)
				info := strings.Split(strings.TrimSpace(str), "-")
				if len(info) != 3 {
					logger.ALogger().Errorf("Novel %s Get Info Error 2,Url %s", novel.NovelName, novel.NovelUrl)
					continue
				}
				novel.NovelType = info[0]
				novel.Status = info[1]
				novel.Desc = info[2]
				noveldb.UpdateOneDataToNovelByNameAndAuthor(novel)
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)
	}

}

func main() {

	args := os.Args
	if args == nil || len(args) < 2 || args[1] == "update" {
		//UpdateData()
		fmt.Println("Hello ", args[1]) // 第二个参数，第一个参数为命令名

	} else if args[1] == "insert" {
		fmt.Println("Hello ", args[1]) // 第二个参数，第一个参数为命令名
		//InsertData()
	} else {
		logger.ALogger().Error("Wrong Input..eg:go run main.go update/insert")
	}

	//UpdateData()
}
