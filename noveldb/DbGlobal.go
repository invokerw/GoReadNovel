package noveldb

import (
	"GoReadNovel/logger"
	"database/sql"
)

var (
	dbhostsip  = "fsnsaber.cn:3306" //IP地址
	dbusername = "root"             //用户名
	dbpassword = "weifei"           //密码
	dbname     = "novel"            //数据库名
)

type Novel struct {
	ID            int    `json:"id"`         //ID
	NovelName     string `json:"novelname"`  //章节名称
	NovelUrl      string `json:"url"`        //地址
	NovelType     string `json:"noveltype"`  //小说类型
	LatestChpName string `json:"newchp"`     //最新章节名字
	LatestChpUrl  string `json:"newurl"`     //最新章节地址
	ImagesAddr    string `json:"imagesaddr"` //封面图片地址
	Author        string `json:"author"`     //作者
	Status        string `json:"status"`     //状态连载还是完结
	Desc          string `json:"desc"`       //描述
}

const (
	DEFAULT_NOVEL_TYPE = "其他"
	DEFAULT_STATUS     = "连载中"
)

var db *sql.DB
var err error

func init() {
	//func mysqlInit() {
	logger.ALogger().Debug("Init Mysql DB Connect..")
	db, err = sql.Open("mysql", "root:weifei@tcp(fsnsaber.cn:3306)/novel?charset=utf8")
	checkErr(err)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(100)
}

func checkErr(err error) {
	if err != nil {
		logger.ALogger().Error(err)
	}
}

func GetMysqlDB() *sql.DB {
	/*if db == nil {
		mysqlInit()
	}*/
	return db
}

func CloseMysqlDB() {
	db.Close()
}
