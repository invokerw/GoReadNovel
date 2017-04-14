package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//增
func InsertOneDataToNovel(novel Novel) {
	//插入数据
	stmt, err := db.Prepare("INSERT novel SET name=?,author=?,noveldesc=?,noveltype=?,addr=?,imageaddr=?,lchaptername=?,lchapteraddr=?,status=?")
	checkErr(err)

	//res, err := stmt.Exec("圣墟", "辰东", "http://www.huanyue123.com/0/11/")
	_, err = stmt.Exec(novel.NovelName, novel.Author, novel.Desc, novel.NovelType, novel.NovelUrl, novel.ImagesAddr,
		novel.LatestChpName, novel.LatestChpUrl, novel.Status)
	checkErr(err)

	//id, err := res.LastInsertId()
	//checkErr(err)

	//logger.ALogger().Debugf("Novel : %v", novel)
}

//改
func UpdateOneDataToNovelByNameAndAuthor(novel Novel) {
	stmt, err := db.Prepare("update novel set noveldesc=?,noveltype=?,lchaptername=?,lchapteraddr=?,status=? where name=? AND author=?")
	checkErr(err)

	res, err := stmt.Exec(novel.Desc, novel.NovelType, novel.LatestChpName, novel.LatestChpUrl, novel.Status, novel.NovelName, novel.Author)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	//logger.ALogger().Debug("updata ,", affect)
}

//查
func FindOneDataToNovelByNovelNameAndAuthor(novel Novel) (Novel, bool) {

	row := db.QueryRow("SELECT * FROM novel WHERE name=? AND author=?", novel.NovelName, novel.Author)
	checkErr(err)
	var novel Novel

	err = row.Scan(&novel.ID, &novel.NovelName, &novel.Author, &novel.Desc, &novel.NovelType, &novel.NovelUrl, &novel.ImagesAddr,
		&novel.LatestChpName, &novel.LatestChpUrl, &novel.Status)
	//checkErr(err)
	if err == sql.ErrNoRows {
		checkErr(err)
		return novel, false
	} else if err != nil {
		checkErr(err)
		return novel, false
	}
	logger.ALogger().Debug("Find One Novel: %v\n", novel)
	return novel, true

}

//删
func DeleteOneDataToNovelByName(id int) {

	stmt, err := db.Prepare("delete from novel where novelid=?")
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
	if db == nil {
		fmt.Print("1\n")
	} else if db != nil {
		fmt.Print("2\n")
	}

}
*/
