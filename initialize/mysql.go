package initialize

import (
	"database/sql"
	"gin_mall/global"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func initMysql() {
	var err error
	global.MDB, err = sql.Open("mysql",global.Config.MysqlServer.Address)
	if err != nil {
		panic(err)
	}
	if err= global.MDB.Ping();err!=nil{panic(err)}
	global.MDB.SetConnMaxLifetime(time.Minute * 3)
	global.MDB.SetMaxOpenConns(100)
	global.MDB.SetMaxIdleConns(10)
}