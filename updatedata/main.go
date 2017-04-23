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

var (
	MAX_PAGE     = 344
	MAX_NOVEL    = 10310
	THREAD_NUM   = 2 //这只是个除数，具体多少线程还是要看除出来几个咯
	ALL_VOTE_NUM = 100
	GOOD_NUM_NUM = 100
)

//从开始到结束  (begin,end]
func InsertData(begin int, end int, ch chan string) {
	logger.ALogger().Debugf("Try to Insert Novel Data (%d,%d]", begin, end)

	for page := begin + 1; page <= end; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTypeNovelList.py", "quanbu", "allvisit", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)

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
			novel.Desc = strings.TrimSpace(idUrlName[5])
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
	ch <- fmt.Sprintf("InsertData From %d To %d Is OK At %s\n", begin, end, time.Now().String())
	return
}

//从开始到结束  (begin,end]
func UpdateData(begin int, end int, ch chan string) {
	logger.ALogger().Debugf("Try to Update Novel Data (%d,%d]", begin, end)
	for page := begin + 1; page <= end; page++ {
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
			novel.Desc = strings.TrimSpace(idUrlName[5])
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			//存在novel就更新数据 不存在就插入一条新数据
			if no, exit := noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
			} else {
				//对比一下时间戳如果是12小时内更新过的话那就不更了
				if time.Now().Unix()-no.UpdateTime < int64(60*60*12) {
					continue
				}
				time.Sleep(1 * time.Second)
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
				novel.Desc = strings.TrimSpace(info[2])
				noveldb.UpdateOneDataToNovelByNameAndAuthor(novel)
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)
	}
	ch <- fmt.Sprintf("UpdateData From %d To %d Is OK At %s\n", begin, end, time.Now().String())
	return
}

//插入AllVote表数据
//满足(end-begin)*30 > maxNum ，maxNum更新多少条数据
func InsertAllVoteData(begin int, end int, maxNum int) {

	logger.ALogger().Debugf("Try to Insert AllVote Data (%d,%d]", begin, end)
	index := 1
	for page := begin + 1; page <= end; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTypeNovelList.py", "quanbu", "allvote", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)

		datas := strings.Split(strings.TrimSpace(str), ",")

		for _, data := range datas {
			idUrlName := strings.Split(strings.TrimSpace(data), "--")
			//logger.ALogger().Debug("--------------", len(idUrlName))
			if len(idUrlName) != 9 {
				continue
			}
			novel := noveldb.Novel{}
			novel.NovelUrl = idUrlName[2] //"/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
			novel.NovelName = idUrlName[3]
			novel.LatestChpName = idUrlName[7]
			novel.Author = idUrlName[4]
			novel.Desc = strings.TrimSpace(idUrlName[5])
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			//查询有没有，没有的话插入小说，再更新
			if no, exit := noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
				if no, exit = noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
					allVote := noveldb.AllVote{}
					allVote.NovelID = no.ID
					index = index + 1
					noveldb.InsertOneDataToAllVote(allVote)
				}
			} else {
				allVote := noveldb.AllVote{}
				allVote.NovelID = no.ID
				index = index + 1
				noveldb.InsertOneDataToAllVote(allVote)
			}
			if maxNum == index {
				logger.ALogger().Debugf("Insert %d Novel By AllVote Over.", maxNum)
				return
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)

	}
}

//更新AllVote表数据
//满足(end-begin)*30 > maxNum ，maxNum更新多少条数据
func UpdateAllvoteData(begin int, end int, maxNum int) {
	logger.ALogger().Debugf("Try to Update AllVote Data (%d,%d]", begin, end)

	index := 1
	for page := begin + 1; page <= end; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTyGpeNovelList.py", "quanbu", "allvote", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)
		datas := strings.Split(strings.TrimSpace(str), ",")

		for _, data := range datas {
			idUrlName := strings.Split(strings.TrimSpace(data), "--")
			//logger.ALogger().Debug("--------------", len(idUrlName))
			if len(idUrlName) != 9 {
				continue
			}

			novel := noveldb.Novel{}
			novel.NovelUrl = idUrlName[2] //"/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
			novel.NovelName = idUrlName[3]
			novel.LatestChpName = idUrlName[7]
			novel.Author = idUrlName[4]
			novel.Desc = strings.TrimSpace(idUrlName[5])
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			//存在novel就更新数据 不存在就插入一条新数据
			if no, exit := noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
				if no, exit = noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
					allVote := noveldb.AllVote{}
					allVote.AllVoteID = index
					index = index + 1
					allVote.NovelID = no.ID
					noveldb.InsertOneDataToAllVote(allVote)
				}
			} else {
				allVote := noveldb.AllVote{}
				allVote.AllVoteID = index
				index = index + 1
				allVote.NovelID = no.ID
				noveldb.InsertOneDataToAllVote(allVote)
			}
			if maxNum == index {
				logger.ALogger().Debugf("Update %d Novel By AllVote Over.", maxNum)
				return
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)
	}
	return
}

//插入GoodNum表数据
//满足(end-begin)*30 > maxNum ，maxNum更新多少条数据
func InsertGoodNumData(begin int, end int, maxNum int) {

	logger.ALogger().Debugf("Try to Insert GoodNum Data (%d,%d]", begin, end)
	index := 1
	for page := begin + 1; page <= end; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTypeNovelList.py", "quanbu", "goodnum", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)

		datas := strings.Split(strings.TrimSpace(str), ",")

		for _, data := range datas {
			idUrlName := strings.Split(strings.TrimSpace(data), "--")
			//logger.ALogger().Debug("--------------", len(idUrlName))
			if len(idUrlName) != 9 {
				continue
			}
			novel := noveldb.Novel{}
			novel.NovelUrl = idUrlName[2] //"/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
			novel.NovelName = idUrlName[3]
			novel.LatestChpName = idUrlName[7]
			novel.Author = idUrlName[4]
			novel.Desc = strings.TrimSpace(idUrlName[5])
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			//查询有没有，没有的话插入小说，再更新
			if no, exit := noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
				if no, exit = noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
					goodNum := noveldb.GoodNum{}
					goodNum.NovelID = no.ID
					index = index + 1
					noveldb.InsertOneDataToGoodNum(goodNum)
				}
			} else {
				goodNum := noveldb.GoodNum{}
				goodNum.NovelID = no.ID
				index = index + 1
				noveldb.InsertOneDataToGoodNum(goodNum)
			}
			if maxNum == index {
				logger.ALogger().Debugf("Insert %d Novel By GoodNum Over.", maxNum)
				return
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)

	}
}

//更新GoodNum表数据
//满足(end-begin)*30 > maxNum ，maxNum更新多少条数据
func UpdateGoodNumData(begin int, end int, maxNum int) {
	logger.ALogger().Debugf("Try to Update GoodNum Data (%d,%d]", begin, end)

	index := 1
	for page := begin + 1; page <= end; page++ {
		strPage := strconv.Itoa(page)
		cmd := exec.Command("python", "../python/getTopByTypeNovelList.py", "quanbu", "goodnum", strPage)

		buf, err := cmd.Output()
		if err != nil {
			logger.ALogger().Errorf("Page %d,%v", page, err)
			continue
		}
		str := string(buf)
		datas := strings.Split(strings.TrimSpace(str), ",")

		for _, data := range datas {
			idUrlName := strings.Split(strings.TrimSpace(data), "--")
			//logger.ALogger().Debug("--------------", len(idUrlName))
			if len(idUrlName) != 9 {
				continue
			}

			novel := noveldb.Novel{}
			novel.NovelUrl = idUrlName[2] //"/GetBookInfo?go=" + idUrlName[2][len(URL):len(idUrlName[2])]
			novel.NovelName = idUrlName[3]
			novel.LatestChpName = idUrlName[7]
			novel.Author = idUrlName[4]
			novel.Desc = strings.TrimSpace(idUrlName[5])
			novel.LatestChpUrl = idUrlName[6] //"/GetBookInfo?go=" + idUrlName[6]
			novel.ImagesAddr = idUrlName[8]
			novel.NovelType = noveldb.DEFAULT_NOVEL_TYPE
			novel.Status = noveldb.DEFAULT_STATUS

			//logger.ALogger().Info("Novle:", novel)
			//存在novel就更新数据 不存在就插入一条新数据
			if no, exit := noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
				noveldb.InsertOneDataToNovel(novel)
				if no, exit = noveldb.FindOneDataFromNovelByNameAndAuthor(novel); exit == false {
					goodNum := noveldb.GoodNum{}
					goodNum.GoodNumID = index
					index = index + 1
					goodNum.NovelID = no.ID
					noveldb.UpdateOneDataToGoodNumByGoodNumID(goodNum)
				}
			} else {
				goodNum := noveldb.GoodNum{}
				goodNum.GoodNumID = index
				index = index + 1
				goodNum.NovelID = no.ID
				noveldb.UpdateOneDataToGoodNumByGoodNumID(goodNum)
			}
			if maxNum == index {
				logger.ALogger().Debugf("Update %d Novel By GoodNum Over.", maxNum)
				return
			}
		}
		logger.ALogger().Debugf("Page/All:%d/%d. Sleep 4s", page, MAX_PAGE)
		time.Sleep(4 * time.Second)
	}
	return
}
func main() {

	var ch = make(chan string, THREAD_NUM)
	runUpdateOrInsert := 0
	funcs := map[int]func(int, int, chan string){
		0: UpdateData,
		1: InsertData,
	}
	funcsAllVote := map[int]func(int, int, int){
		0: UpdateAllvoteData,
		1: InsertAllVoteData,
	}
	funcsGoodNum := map[int]func(int, int, int){
		0: UpdateGoodNumData,
		1: InsertGoodNumData,
	}

	args := os.Args
	if args == nil || len(args) < 2 || args[1] == "update" {
		logger.ALogger().Infof("You Run Update")
		runUpdateOrInsert = 0
		//UpdateData()
	} else if args[1] == "insert" {
		logger.ALogger().Infof("You Run Insert")
		runUpdateOrInsert = 1
		//InsertData()
	} else {
		logger.ALogger().Error("Wrong Input..eg:go run main.go update/insert")
		return
	}

	cmd := exec.Command("python", "../python/getMaxPageNum.py")
	buf, err := cmd.Output()
	if err != nil {
		logger.ALogger().Errorf("Main Get MaxPageNum Error %v", err)
		return
	}
	str := string(buf)
	pageAndNovelNum := strings.Split(strings.TrimSpace(str), "-")
	if len(pageAndNovelNum) != 2 {
		logger.ALogger().Errorf("Main Get MaxPageNum Error:Output Wrong ->%s\n", pageAndNovelNum)
		return
	}
	pageNum, err := strconv.Atoi(pageAndNovelNum[0])
	novelNum, err := strconv.Atoi(pageAndNovelNum[1])
	if err != nil {
		logger.ALogger().Errorf("Main Get MaxPageNum Error:A to i Wrong \n")
		return
	}
	if novelNum/30 != pageNum-1 {
		logger.ALogger().Errorf("Main Get MaxPageNum Error:Output Wrong ->%s\n", pageAndNovelNum)
		return
	}
	MAX_PAGE = pageNum
	MAX_NOVEL = novelNum
	logger.ALogger().Debugf("PageNum = %d, NovelNum = %d", MAX_PAGE, MAX_NOVEL)
	thread_num := 0
	for num := 0; num < MAX_PAGE; num = num + MAX_PAGE/THREAD_NUM {
		//num -- num + MAX_PAGE/10
		if num+MAX_PAGE/(THREAD_NUM-1) >= MAX_PAGE {
			logger.ALogger().Debugf("min-max:%d/%d", num, MAX_PAGE)
			thread_num = thread_num + 1
			go funcs[runUpdateOrInsert](num, MAX_PAGE, ch)
		} else {
			logger.ALogger().Debugf("min-max:%d/%d", num, num+MAX_PAGE/THREAD_NUM)
			thread_num = thread_num + 1
			go funcs[runUpdateOrInsert](num, num+MAX_PAGE/THREAD_NUM, ch)
		}

	}

	logger.ALogger().Debugf("Thread_Num = %d", thread_num)
	for i := 0; i < thread_num; i++ {
		time.Sleep(time.Second * 3)
		logger.ALogger().Debugf(<-ch)
	}
	//插入或者更新allVote、GoodNum表  //需要单独跑一次
	funcsAllVote[runUpdateOrInsert](0, 5, ALL_VOTE_NUM)
	funcsGoodNum[runUpdateOrInsert](0, 5, GOOD_NUM_NUM)

}
