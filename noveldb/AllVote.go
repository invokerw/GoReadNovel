package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToAllVote(allVote AllVote) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	allVote.UpdateTime = time.Now().Unix()
	stmt, err := db.Prepare("INSERT allvote SET novelid=?,updatetime=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(allVote.NovelID, allVote.UpdateTime)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert allvote error %v \n", allVote)
	}

	//logger.ALogger().Debugf("allVote : %v", allVote)
}

//改
func UpdateOneDataToAllVoteByAllVoteID(allVote AllVote) {
	//没有更新 join time
	stmt, err := db.Prepare("update allvote set novelid=?,updatetime=? where allvoteid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(allVote.NovelID, allVote.UpdateTime, allVote.AllVoteID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata allvote error %v \n", allVote)
	}
	//logger.ALogger().Debugf("updata allVote %v \n", allVote)
}

//查
func FindOneDataFromAllVoteByAllVoteID(av AllVote) (AllVote, bool) {

	row := db.QueryRow("SELECT * FROM allvote WHERE allvoteid=?", av.AllVoteID)
	//defer rows.Close()如果是读取很多行的话要关闭
	var allVote AllVote

	err = row.Scan(&allVote.AllVoteID, &allVote.NovelID, &allVote.UpdateTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return allVote, false
	} else if err != nil {
		//checkErr(err)
		return allVote, false
	}
	//logger.ALogger().Debugf("Find One allVote: %v\n", allVote)
	return allVote, true

}

//从begin 开始 num条数据  eg:0,1000 1-1000  查询1000条数据
func FindDatasFromAllVote(begin int, num int) (map[int]AllVote, bool) {

	rows, err := db.Query("SELECT * FROM allvote LIMIT ?,?", begin, num)
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var allVotes map[int]AllVote
	number := 0
	allVotes = make(map[int]AllVote)

	for rows.Next() {
		var allVote AllVote
		rows.Scan(&allVote.AllVoteID, &allVote.NovelID, &allVote.UpdateTime)
		allVotes[number] = allVote
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d allVote: %v\n", num, allVotes)
	return allVotes, true
}
