package handler

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 这个意思是要执行 mysql 该包下的文件里所有init()函数，但不会打包这个包，所以无法通过包名来调用包中的其他函数
	"log"
	"time"
)

const SECTION = "database"

var DB *sql.DB

func CreateMySqlConnection() *sql.DB {
	var err error
	//configMap := common.GetConfigMap(SECTION)
	//name := configMap["name"]
	//user := configMap["user"]
	//password := configMap["passwd"]
	//host := configMap["host"]
	//port := configMap["port"]
	//DB, err = sql.Open(
	//	"mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&sql_mode=''", user, password, host, port, name))
	DB, err = sql.Open(
		"mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&sql_mode=''", "root", "chj981005", "111.230.154.86", "3306", "AutoAnalysis"))
	DB.SetConnMaxLifetime(time.Second * 600)
	DB.SetMaxIdleConns(25)
	DB.SetMaxOpenConns(50)
	if err != nil {
		log.Print(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Print(err)
	}
	return DB
}
