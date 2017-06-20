package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToFeedback(feedback Feedback) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	feedback.AddTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("INSERT feedback SET userid=?,feedbacktype=?,content=?,contactinfo=?,solve=?,addtime=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(feedback.UserID, feedback.FeedbackType, feedback.Content, feedback.ContactInfo, feedback.Solve, feedback.AddTime)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert feedback error %v \n", feedback)
	}

	//logger.ALogger().Debugf("allVote : %v", allVote)
}

//改
func UpdateOneDataSolvedToFeedbackByFeedbackID(feedbackid int) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update feedback set solve=? where feedbackid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(1, feedbackid)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata feedback solve error %v \n", feedbackid)
	}
	//logger.ALogger().Debugf("updata allVote %v \n", allVote)
}

//查
func FindOneDataFromFeedbackByFeedbackID(feedbackid int) (Feedback, bool) {

	row := GetMysqlDB().QueryRow("SELECT * FROM feedback WHERE feedbackid=?", feedbackid)
	//defer rows.Close()如果是读取很多行的话要关闭
	var feedback Feedback

	err = row.Scan(&feedback.FeedbackID, &feedback.UserID, &feedback.FeedbackType, &feedback.Content,
		&feedback.ContactInfo, &feedback.Solve, &feedback.AddTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return feedback, false
	} else if err != nil {
		//checkErr(err)
		return feedback, false
	}
	//logger.ALogger().Debugf("Find One feedback: %v\n", feedback)
	return feedback, true

}

//查询的问题  feedback类型 -1 为所有 sovel 0为未解决 1为已解决 -1为所有
func FindDatasFromFeedback(feedbacktype int, solved int) (map[int]Feedback, bool) {
	var rows *sql.Rows
	var err error
	if feedbacktype == -1 && solved == -1 {
		rows, err = GetMysqlDB().Query("SELECT * FROM feedback")
	} else if feedbacktype == -1 && solved != -1 {
		rows, err = GetMysqlDB().Query("SELECT * FROM feedback where solve=?",
			solved)
	} else if feedbacktype != -1 && solved == -1 {
		rows, err = GetMysqlDB().Query("SELECT * FROM feedback where feedbacktype=? ",
			feedbacktype)
	} else {
		rows, err = GetMysqlDB().Query("SELECT * FROM feedback where feedbacktype=? and solve=?",
			feedbacktype, solved)
	}
	defer rows.Close() //如果是读取很多行的话要关闭

	if !checkErr(err) {
		return nil, false
	}

	var feedbacks map[int]Feedback
	number := 0
	feedbacks = make(map[int]Feedback)

	for rows.Next() {
		var feedback Feedback
		rows.Scan(&feedback.FeedbackID, &feedback.UserID, &feedback.FeedbackType, &feedback.Content,
			&feedback.ContactInfo, &feedback.Solve, &feedback.AddTime)
		feedbacks[number] = feedback
		number = number + 1
	}
	//logger.ALogger().Debugf("Find %d feedback: %v\n", number, feedbacks)
	return feedbacks, true
}

//删
func DeleteOneDataToFeedbackByID(feedbackid int) bool {

	stmt, err := GetMysqlDB().Prepare("delete from feedback where feedbackid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(feedbackid)
	if !checkErr(err) {
		return false
	}
	return true
}
