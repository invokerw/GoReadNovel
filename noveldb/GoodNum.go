package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToGoodNum(goodNum GoodNum) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	goodNum.UpdateTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("INSERT goodnum SET novelid=?,updatetime=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(goodNum.NovelID, goodNum.UpdateTime)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert goodNum error %v \n", goodNum)
	}

	//logger.ALogger().Debugf("goodNum : %v", goodNum)
}

//改
func UpdateOneDataToGoodNumByGoodNumID(goodNum GoodNum) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update goodnum set novelid=?,updatetime=? where goodnumid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(goodNum.NovelID, goodNum.UpdateTime, goodNum.GoodNumID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata goodNum error %v \n", goodNum)
	}
	//logger.ALogger().Debugf("updata goodNum %v \n", goodNum)
}

//查
func FindOneDataFromGoodNumByGoodNumID(gn GoodNum) (GoodNum, bool) {

	row := GetMysqlDB().QueryRow("SELECT * FROM goodnum WHERE goodnumid=?", gn.GoodNumID)
	//defer rows.Close()如果是读取很多行的话要关闭
	var goodNum GoodNum

	err = row.Scan(&goodNum.GoodNumID, &goodNum.NovelID, &goodNum.UpdateTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return goodNum, false
	} else if err != nil {
		//checkErr(err)
		return goodNum, false
	}
	//logger.ALogger().Debugf("Find One goodNum: %v\n", goodNum)
	return goodNum, true

}

//从begin 开始 num条数据  eg:0,1000 1-1000  查询1000条数据
func FindDatasFromGoodNum(begin int, num int) (map[int]GoodNum, bool) {

	rows, err := GetMysqlDB().Query("SELECT * FROM goodnum LIMIT?,?", begin, num)
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var goodNums map[int]GoodNum
	number := 0
	goodNums = make(map[int]GoodNum)

	for rows.Next() {
		var goodNum GoodNum
		rows.Scan(&goodNum.GoodNumID, &goodNum.NovelID, &goodNum.UpdateTime)
		goodNums[number] = goodNum
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d goodNums: %v\n", num, goodNums)
	return goodNums, true
}
