package main

// database\mysql 库的使用

// 第一步： 连接数据库
// Go语言中的database/sql包提供了保证SQL或类SQL数据库的泛用接口，并不提供具体的数据库驱动。使用database/sql包时必须注入（至少）一个数据库驱动。

import (
	"database/sql"
	db2 "error_learn/db"
	"error_learn/mysql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB


func main(){

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%+v", err)
		}
	}()

	db,err := db2.InitDB()
	defer db.Close()
	if err != nil{
		fmt.Printf("stack : %+v", err)
		return
	}

	fmt.Println("连接成功")
	 partsModel := mysql.NewPart(db)
	 parts, err := partsModel.GetById(1)
	 if err != nil {
	 	fmt.Printf("%+v","part not found")
		 return
	 }
	 fmt.Printf("%v",parts)
}






