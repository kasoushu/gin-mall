package api

import (
	"fmt"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreteAddress(ctx *gin.Context) {
	var address model.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if address.Province == "" || address.City == "" || address.District == "" {
		model.Failed("params error", ctx)
		return
	}

	if service.InsertAddress(address) {
		model.Success("create successful!", "", ctx)
		return
	}
	model.Failed("create fail", ctx)
}

func UpdateAddress(ctx *gin.Context) {
	var address model.Address

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if err := ctx.ShouldBindJSON(&address); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if service.UpdateAddress(id, address) {
		model.Success("update successful!", "", ctx)
		return
	}
	model.Failed("update fail", ctx)
}
func DeleteAddress(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if service.DeleteAddress(id) {
		model.Success("delete successful!", "", ctx)
		return
	}
	model.Failed("delete fail", ctx)
}

func GetSingleAddressInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	var address model.Address
	if service.GetAddressInfo(id, &address) {
		//p, err := json.Marshal(address)
		if err != nil {
			fmt.Println(err)
			model.Failed("convert json error", ctx)
		}
		model.Success("get info successful!", address, ctx)
		return
	}

	model.Failed("get info fail", ctx)

}
