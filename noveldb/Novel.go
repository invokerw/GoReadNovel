package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToNovel(novel Novel) {
	//插入数据
	novel.UpdateTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("INSERT novel SET name=?,author=?,noveldesc=?,noveltype=?,addr=?,imageaddr=?,lchaptername=?,lchapteraddr=?,status=?,updatetime=?")
	defer stmt.Close()
	checkErr(err)

	//res, err := stmt.Exec("圣墟", "辰东", "http://www.huanyue123.com/0/11/")
	_, err = stmt.Exec(novel.NovelName, novel.Author, novel.Desc, novel.NovelType, novel.NovelUrl, novel.ImagesAddr,
		novel.LatestChpName, novel.LatestChpUrl, novel.Status, novel.UpdateTime)
	if !checkErr(err) {
		logger.ALogger().Errorf("insert novel error %v \n", novel)
	}

	//id, err := res.LastInsertId()
	//checkErr(err)

	//logger.ALogger().Debugf("Novel : %v", novel)
}

//改
func UpdateOneDataToNovelByNameAndAuthor(novel Novel) {

	novel.UpdateTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("update novel set noveldesc=?,noveltype=?,addr=?,imageaddr=?,lchaptername=?,lchapteraddr=?,status=?,updatetime=? where name=? AND author=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(novel.Desc, novel.NovelType, novel.NovelUrl, novel.ImagesAddr, novel.LatestChpName, novel.LatestChpUrl,
		novel.Status, novel.UpdateTime, novel.NovelName, novel.Author)
	if !checkErr(err) {
		logger.ALogger().Debugf("updata novel error %v \n", novel)
	}

	//affect, err := res.RowsAffected()
	//checkErr(err)
	//logger.ALogger().Debugf("updata novel %v \n", novel)
}

//查询特定的一条数据依据小说名称与作者
func FindOneDataFromNovelByNameAndAuthor(no Novel) (Novel, bool) {

	row := GetMysqlDB().QueryRow("SELECT * FROM novel WHERE name=? AND author=?", no.NovelName, no.Author)
	//defer rows.Close()如果是读取很多行的话要关闭

	var novel Novel

	err = row.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
		&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status, &novel.UpdateTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return novel, false
	} else if err != nil {
		//checkErr(err)
		return novel, false
	}
	//logger.ALogger().Debugf("Find One Novel: %v\n", novel)
	return novel, true

}

//查询特定的一条数据依据小说NovelID
func FindOneDataFromNovelByID(novelId int) (Novel, bool) {

	row := GetMysqlDB().QueryRow("SELECT * FROM novel WHERE novelid=?", novelId)
	//defer rows.Close()如果是读取很多行的话要关闭

	var novel Novel

	err = row.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
		&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status, &novel.UpdateTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return novel, false
	} else if err != nil {
		//checkErr(err)
		return novel, false
	}
	//logger.ALogger().Debugf("Find One Novel: %v\n", novel)
	return novel, true

}

//从begin 开始 num条数据  eg:0,1000 1-1000  查询1000条数据
func FindDatasFromNovel(begin int, num int) (map[int]Novel, bool) {

	rows, err := GetMysqlDB().Query("SELECT * FROM novel LIMIT?,?", begin, num)
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var novels map[int]Novel
	number := 0
	novels = make(map[int]Novel)

	for rows.Next() {
		var novel Novel
		rows.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
			&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status, &novel.UpdateTime)
		novels[number] = novel
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d novels: %v\n", num, novels)
	return novels, true
}

//查询若干条数据依据模糊的小说名称或者某个作者
func FindDatasFromNovelByNameOrAuthor(key string) (map[int]Novel, bool) {
	rows, err := GetMysqlDB().Query("SELECT * FROM novel WHERE name LIKE '%?% or' author LIKE '%?%'", key, key)
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var novels map[int]Novel
	number := 0
	novels = make(map[int]Novel)

	for rows.Next() {
		var novel Novel
		rows.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
			&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status, &novel.UpdateTime)
		novels[number] = novel
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d novels: %v\n", num, novels)
	return novels, true
}

//查询若干条数据依据小说类型 没限制数量
func FindDatasFromNovelByNovelType(novelType string) (map[int]Novel, bool) {
	rows, err := GetMysqlDB().Query("SELECT * FROM novel WHERE noveltype=?", novelType)
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var novels map[int]Novel
	number := 0
	novels = make(map[int]Novel)

	for rows.Next() {
		var novel Novel
		rows.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
			&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status, &novel.UpdateTime)
		novels[number] = novel
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d novels: %v\n", num, novels)
	return novels, true
}

//删
func DeleteOneDataToNovelByName(id int) {

	stmt, err := GetMysqlDB().Prepare("delete from novel where novelid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	logger.ALogger().Debug(affect)
}

/*
func main() {

	fmt.Println("vim-go")
	if GetMysqlDB() == nil {
		fmt.Print("1\n")
	} else if GetMysqlDB() != nil {
		fmt.Print("2\n")
	}

}
*/
