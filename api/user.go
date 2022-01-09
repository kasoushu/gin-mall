package api

import (
	"fmt"
	"gin_mall/common"
	"gin_mall/global"
	"gin_mall/middleware"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context)  {
	var uu model.User
	var flag int =-1
	if err :=c.ShouldBind(&uu);err!=nil{
		model.Failed("information input error",c)
		return
	}
	if len(uu.Phone)!=11{
		model.Failed("phone number error",c)
		return
	}
	row := global.MDB.QueryRow("select id from users where phone=?",uu.Phone)
	if row.Err()!=nil {
		fmt.Println(row)
		model.Failed("query error",c)
		return
	}
	if row.Scan(&flag);flag==-1{
		common.InsertUser(uu)
		model.Success("signup successful!",true,c)
		return
	}
	model.Failed("phone had existed",c)
}

func UserLogin(c *gin.Context){
	var p model.WebLoginUser
	if err :=c.ShouldBind(&p);err!=nil{
		model.Failed("密码或手机号出错",c)
	}
	var uu model.User
	ok := service.IsExist(p,&uu)
	//fmt.Println(p.Phone,p.Password,id)
	if ok{
		token,err := middleware.CreateToken(uu.Name,uu.Phone)
		if err != nil {
			model.Failed("token error",c)
		}
		model.Success("successful login ",gin.H{
			"id":uu.Id,
			"name":uu.Name,
			"age":uu.Age,
			"token":token,
		},c)
		return
	}
	model.Failed("密码或手机号出错",c)

}