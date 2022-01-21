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

func AdminSIgnUp(c *gin.Context) {
	var uu model.Admin
	var flag int = -1
	if err := c.ShouldBind(&uu); err != nil {
		model.Failed("information input error", c)
		return
	}
	if len(uu.Phone) != 11 {
		model.Failed("phone number error", c)
		return
	}
	row := global.MDB.QueryRow("select id from admins where phone=?", uu.Phone)
	if row.Err() != nil {
		fmt.Println(row)
		model.Failed("query error", c)
		return
	}
	if row.Scan(&flag); flag == -1 {
		_, err := global.MDB.Exec("insert into admins(name,phone,password) values (?,?,?)",
			uu.Name, uu.Phone, uu.Password)
		if err != nil {
			fmt.Println(err)
			model.Failed("insert fail", c)
			return
		}
		model.Success("signup successful!", true, c)
		return
	}
	model.Failed("phone had existed", c)
}

func UserSignUp(c *gin.Context) {
	var uu model.User
	var flag int = -1
	if err := c.ShouldBind(&uu); err != nil {
		model.Failed("information input error", c)
		return
	}
	if len(uu.Phone) != 11 {
		model.Failed("phone number error", c)
		return
	}
	row := global.MDB.QueryRow("select id from users where phone=?", uu.Phone)
	if row.Err() != nil {
		fmt.Println(row)
		model.Failed("query error", c)
		return
	}
	if row.Scan(&flag); flag == -1 {
		if common.InsertUser(uu) {
			model.Success("signup successful!", true, c)
			return
		}
		fmt.Println("insert error")
		model.Failed("insert error", c)
		return
	}
	model.Failed("phone had existed", c)
}

func UserLogin(c *gin.Context) {
	var p model.WebLoginUser
	if err := c.ShouldBind(&p); err != nil {
		model.Failed("密码或手机号出错", c)
	}
	var uu model.User
	ok := service.IsExist(p, &uu)
	//fmt.Println(p.Phone,p.Password,id)
	if ok {
		token, err := middleware.CreateToken(uint64(uu.Id), uu.Phone)
		if err != nil {
			model.Failed("token error", c)
		}
		model.Success("successful login ", gin.H{
			"id":    uu.Id,
			"name":  uu.Name,
			"token": token,
			"phone": uu.Phone,
		}, c)
		return
	}
	model.Failed("密码或手机号出错", c)
}
func AdminLogin(c *gin.Context) {
	var p model.Admin
	if err := c.ShouldBind(&p); err != nil {
		model.Failed("密码或手机号出错", c)
	}
	var uu model.Admin
	row := global.MDB.QueryRow("select id, name,password,phone from admins where admins.phone=? ",
		p.Phone)
	if row.Err() != nil {
		fmt.Println(row.Err())
		model.Failed("error 账户不存在", c)
		return
	}
	if err := row.Scan(&uu.Id, &uu.Name, &uu.Password, &uu.Phone); err != nil {
		fmt.Println(err)
		model.Failed("error", c)
		return
	}
	if uu.Password == p.Password {
		token, err := middleware.CreateToken(uu.Id, uu.Phone)
		if err != nil {
			model.Failed("token error", c)
		}
		model.Success("successful login ", gin.H{
			"id":    uu.Id,
			"name":  uu.Name,
			"phone": uu.Phone,
			"token": token,
		}, c)
		return
	}
	model.Failed("密码或手机号出错", c)
}
