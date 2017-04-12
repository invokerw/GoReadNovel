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

func main() {
	db, err := sql.Open("mysql", "root:weifei@tcp(fsnsaber.cn:3306)/novel?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT novel SET name=?,author=?,addr=?")
	checkErr(err)

	res, err := stmt.Exec("圣墟", "辰东", "http://www.huanyue123.com/0/11/")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update novel set name=? where novelid=?")
	checkErr(err)

	res, err = stmt.Exec("圣墟2", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM novel")
	checkErr(err)

	for rows.Next() {
		var novelid int
		var name string
		var author string
		var addr string
		var aaa string
		err = rows.Scan(&novelid, &name, &author, &aaa, &aaa, &addr, &aaa, &aaa, &aaa, &aaa, &aaa)
		//checkErr(err)
		fmt.Println(novelid)
		fmt.Println(name)
		fmt.Println(author)
		fmt.Println(addr)
	}

	//删除数据
	stmt, err = db.Prepare("delete from novel where novelid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		fmt.Print(err)
	}
}
