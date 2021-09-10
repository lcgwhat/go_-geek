package db

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)
type mysqlConfig struct {
	username string
	password string
	host string
	port int
	dataName string
}
func InitDB()(*sql.DB,error){
	var mysqlConfig = mysqlConfig{
		username: "root",
		password: "",
		host: "127.0.0.1",
		port: 3306,
		dataName: "pms_dev",
	}
	// 初始化连接
	var dns = ""
	var err error
	dns = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", mysqlConfig.username, mysqlConfig.password, mysqlConfig.host, mysqlConfig.port, mysqlConfig.dataName)
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		panic(errors.Wrap(err, "数据库连接失败"))
	}
	return db,nil
}