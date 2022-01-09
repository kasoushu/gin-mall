package common

import (
"database/sql"
"gin_mall/global"
"gin_mall/model"
)

func QueryUser(u model.WebLoginUser)(*sql.Row,error){
	stm ,err := global.MDB.Prepare("select id,name,age,phone from users where phone=? and password=?")
	if err != nil {
		return nil,err
	}
	row := stm.QueryRow(u.Phone,u.Password)
	return row,nil
}

func InsertUser(user model.User) bool{
	stm,err := global.MDB.Prepare("insert into users(name,age,phone,password) values (?,?,?,?)")
	if err != nil {
		return false
	}
	_,err= stm.Exec(user.Name,user.Age,user.Phone,user.Password)
	if err!=nil {
		return false
	}
	return true
}

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
