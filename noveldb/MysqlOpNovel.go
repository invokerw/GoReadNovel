package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "fsnsaber.cn:3306" //IP地址
	dbusername = "root"             //用户名
	dbpassword = "weifei"           //密码
	dbname     = "novel"            //数据库名
)

type Novel struct {
	ID            int    `json:"id"`        //ID
	NovelName     string `json:"novelname"` //章节名称
	NovelUrl      string `json:"url"`       //地址
	LatestChpName string `json:"newchp"`    //最新章节名字
	LatestChpUrl  string `json:"newurl"`    //最新章节地址
	ImagesAddr    string `json:"imagesaddr"`//封面图片地址
	Author        string `json:"author"`    //作者
	Status        string `json:"status"`    //状态连载还是完结
	Desc          string `json:"desc"`      //描述
}

var db *sql.DB
var err error

func init() {
	//func mysqlInit() {
	fmt.Println("Init Mysql DB Connect..")
	db, err = sql.Open("mysql", "root:weifei@tcp(fsnsaber.cn:3306)/novel?charset=utf8")
	checkErr(err)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(100)
}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func GetMysqlDB() *sql.DB {
	/*if db == nil {
		mysqlInit()
	}*/
	return db
}

//增
func InsertDataToNovel() {
	
}

//改
func UpdateDataToNovel() {
	
}
//查
func FindDataToNovel() {
	
}
//删
func DeleteDataToNovel() {
	
}



func main() {
	fmt.Println("vim-go")
	if db == nil {
		fmt.Print("1\n")
	} else if db != nil {
		fmt.Print("2\n")
	}
}
