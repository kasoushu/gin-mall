package service

import (
	"gin_mall/common"
	"gin_mall/model"
)




func IsExist(w model.WebLoginUser,u *model.User) bool{
	row,err:=common.QueryUser(w)
	if err != nil {
		return false
	}
	if row!=nil{
		err = row.Scan(&u.Id,&u.Name,&u.Age,&u.Phone)
		if err!=nil{return false}
		return true
	}
	return false
}