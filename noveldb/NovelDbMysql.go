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

func main() {
	fmt.Println("vim-go")
	if db == nil {
		fmt.Print("1\n")
	} else if db != nil {
		fmt.Print("2\n")
	}
}
