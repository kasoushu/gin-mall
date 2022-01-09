package main

import (
	"gin_mall/global"
	"gin_mall/initialize"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	name string
	age int
	phone string
	password string
}



func main() {
	r := gin.Default()
	initialize.Initial()
	add :=global.Config.Server.Address+":"+global.Config.Server.Port
	defer global.MDB.Close()
	//stt ,_ := db.Prepare("insert into users(name,age,phone,password) values (?,?,?,?)")
	//_,err =stt.Exec("lalla",20,"18888888888","0000")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//sqt,_:= db.Prepare("select * from users where phone='10000000000'")
	//rows,_:=sqt.Query()

	//var user User
	//rows,_ := global.MDB.Query("select * from  users ")
	//for rows.Next() {
	//	e := rows.Scan(&user.id,&user.name,&user.age,&user.phone,&user.password)
	//	fmt.Println(user.id,"  ",user.name," ",user.phone)
	//	if e!=nil {
	//		fmt.Println(e)
	//		break
	//	}
	//}
	//rows.Close()
	r.Run(add)

}
