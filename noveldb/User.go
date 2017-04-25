package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//增
func InsertOneDataToUser(user User) {
	//插入数据  插入数据时候不需要写入时间，插入时候会帮助你写入
	user.JoinTime = time.Now().Unix()
	stmt, err := GetMysqlDB().Prepare("INSERT user SET userid=?,nikename=?,gender=?,city=?,province=?,country=?,avatarurl=?,jointime=?")
	defer stmt.Close()
	checkErr(err)

	//res, err := stmt.Exec("圣墟", "辰东", "http://www.huanyue123.com/0/11/")
	_, err = stmt.Exec(user.UserID, user.NikeName, user.Gender, user.City, user.Province, user.Country, user.AvatarUrl, user.JoinTime)

	if !checkErr(err) {
		logger.ALogger().Errorf("insert user error %v \n", user)
	}

	//logger.ALogger().Debugf("user : %v", user)
}

//改
func UpdateOneDataToUserByUserID(user User) {
	//没有更新 join time
	stmt, err := GetMysqlDB().Prepare("update user set nikename=?,gender=?,city=?,province=?,country=?,avatarurl=? where userid=?")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(user.NikeName, user.Gender, user.City, user.Province, user.Country, user.AvatarUrl, user.UserID)
	if !checkErr(err) {
		logger.ALogger().Errorf("updata user error %v \n", user)
	}
	//logger.ALogger().Debugf("updata user %v \n", user)
}

//查
func FindOneDataFromUserByUserID(uid string) (User, bool) {

	row := GetMysqlDB().QueryRow("SELECT * FROM user WHERE userid=?", uid)
	//defer rows.Close()如果是读取很多行的话要关闭
	var user User

	err = row.Scan(&user.UserID, &user.NikeName, &user.Gender, &user.City, &user.Province, &user.Country, &user.AvatarUrl, &user.JoinTime)
	//checkErr(err)
	if err == sql.ErrNoRows {
		//checkErr(err)
		//查不到就不报Error了
		return user, false
	} else if err != nil {
		//checkErr(err)
		return user, false
	}
	//logger.ALogger().Debugf("Find One user: %v\n", user)
	return user, true

}
