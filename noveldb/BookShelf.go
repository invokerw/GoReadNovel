package noveldb

import (
	"GoReadNovel/logger"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//增
func InsertOneDataToBookShelf(bookShelf BookShelf) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	stmt, err := GetMysqlDB().Prepare("INSERT bookshelf SET shelfid=?,userid=?,novelid=?,readchaptername=?,readchapteraddr=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(bookShelf.ShelfID, bookShelf.UserID, bookShelf.NovelID, bookShelf.ReadChapterName, bookShelf.ReadChapterUrl)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert bookshelf error %v \n", bookShelf)
	}

	//logger.ALogger().Debugf("bookShelf : %v", bookShelf)
}

//改
func UpdateOneDataToBookShlefByUserIDAndNovelID(bookShelf BookShelf) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update bookshelf set readchaptername=?,readchapteraddr=? where userid=? and novelid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(bookShelf.ReadChapterName, bookShelf.ReadChapterUrl, bookShelf.UserID, bookShelf.NovelID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata user bookShelf %v \n", bookShelf)
	}
	//logger.ALogger().Debugf("updata bookShelf %v \n", user)
}

//查
func FindOneUserBookShlefFromBookShelfByUserID(userid string) (map[int]BookShelf, bool) {

	rows, err := GetMysqlDB().Query("SELECT * FROM bookshelf WHERE userid=?", userid)
	defer rows.Close() //如果是读取很多行的话要关闭
	var bookShelfs map[int]BookShelf
	bookShelfs = make(map[int]BookShelf)

	num := 0
	if !checkErr(err) {
		//如果出现这个那就是bug了。。
		return bookShelfs, false
	}

	for rows.Next() {
		var bookShelf BookShelf
		rows.Scan(&bookShelf.ShelfID, &bookShelf.UserID, &bookShelf.NovelID, &bookShelf.ReadChapterName, &bookShelf.ReadChapterUrl)
		bookShelfs[num] = bookShelf
		num = num + 1
	}
	//logger.ALogger().Debugf("Find %d bookshlef: %v\n", num,bookShelfs)
	if len(bookShelfs) == 0 {
		return bookShelfs, false
	}
	return bookShelfs, true

}

//删除某一条书架上书籍数据
func DeleteOneDataToBookShelfByUseridAndNovelid(userid string, novelid int) {

	stmt, err := GetMysqlDB().Prepare("delete from novel where novelid=? and userid=?")
	checkErr(err)

	_, err = stmt.Exec(novelid, userid)
	checkErr(err)
	//理论上不会错呀
	//logger.ALogger().Debug(affect)
}
