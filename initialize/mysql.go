package initialize

import (
	"database/sql"
	"gin_mall/global"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func initMysql() {
	var err error
	global.MDB, err = sql.Open("mysql", global.Config.MysqlServer.Address)
	if err != nil {
		panic(err)
	}
	if err = global.MDB.Ping(); err != nil {
		panic(err)
	}
	global.MDB.SetConnMaxLifetime(time.Minute * 3)
	global.MDB.SetMaxOpenConns(100)
	global.MDB.SetMaxIdleConns(10)

	//charu
	//session, err := global.MDB.Begin()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for i := 0; i < 1000; i++ {
	//	_, err := session.Exec(`insert into products(category_id,
	//             price,amount,title,created,updated,
	//             created_by,name ) VALUES (100000002,1,1,"sad",now(),now(),10002,'test'); `)
	//	if err != nil {
	//		fmt.Println(err)
	//		//break
	//
	//	}
	//}
	//_ = session.Commit()
}
