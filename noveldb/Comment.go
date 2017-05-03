package noveldb

import (
	"GoReadNovel/logger"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToComment(comment Comment) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	comment.CommentTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("INSERT comment SET commentid=?,userid=?,novelid=?,content=?,commenttime=?,zan=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(comment.CommentID, comment.UserID, comment.NovelID, comment.Content, comment.CommentTime, comment.Zan)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert comment error %v \n", comment)
	}

	//logger.ALogger().Debugf("comment : %v", comment)
}

//改 点赞数+1
func UpdateOneDataAddZanToCommentByCommentID(commentID int) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update comment set zan=zan+1 where commentid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(commentID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata user commentID zan %d \n", commentID)
	}
	//logger.ALogger().Debugf("updata bookShelf %v \n", user)
}

//减一
func UpdateOneDataMinZanToCommentByCommentID(commentID int) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update comment set zan=zan-1 where commentid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(commentID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata user commentID zan %d \n", commentID)
	}
	//logger.ALogger().Debugf("updata bookShelf %v \n", user)
}

//查某个书籍的所有评论
func FindOneNovelCommentFromCommentByNovelID(novelid int) (map[int]Comment, bool) {

	rows, err := GetMysqlDB().Query("SELECT * FROM comment WHERE novelid=? limit 20", novelid)
	defer rows.Close() //如果是读取很多行的话要关闭
	var comments map[int]Comment
	comments = make(map[int]Comment)

	num := 0
	if !checkErr(err) {
		//如果出现这个那就是bug了。。
		return comments, false
	}

	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.CommentID, &comment.UserID, &comment.NovelID, &comment.Content, &comment.CommentTime, &comment.Zan)
		comments[num] = comment
		num = num + 1
	}
	//logger.ALogger().Debugf("Find %d bookshlef: %v\n", num,bookShelfs)
	if len(comments) == 0 {
		return comments, false
	}
	return comments, true

}

//删除某一条评论数据
func DeleteOneDataToCommentByCommentid(commentid int) {

	stmt, err := GetMysqlDB().Prepare("delete from comment where commentid=?")
	checkErr(err)

	_, err = stmt.Exec(commentid)
	checkErr(err)

	//理论上不会错呀
	//logger.ALogger().Debug(affect)
}
