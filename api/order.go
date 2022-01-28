package api

import (
	"fmt"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreteOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if order.Status == "" {
		model.Failed("params error", ctx)
		return
	}

	if service.InsertOrder(order) {
		model.Success("create successful!", "", ctx)
		return
	}
	model.Failed("create fail", ctx)
}

func UpdateOrder(ctx *gin.Context) {
	var order model.Order

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if err := ctx.ShouldBindJSON(&order); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if service.UpdateOrder(id, order) {
		model.Success("update successful!", "", ctx)
		return
	}
	model.Failed("update fail", ctx)
}
func DeleteOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if service.DeleteOrder(id) {
		model.Success("delete successful!", "", ctx)
		return
	}
	model.Failed("delete fail", ctx)
}

//func GetSingleOrderInfo(ctx *gin.Context) {
//	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
//	if err != nil {
//		fmt.Println(err)
//		model.Failed("parse params error", ctx)
//		return
//	}
//	var order model.Order
//	if service.GetOrderInfo(id, &order) {
//		//p, err := json.Marshal(order)
//		if err != nil {
//			fmt.Println(err)
//			model.Failed("convert json error", ctx)
//		}
//		model.Success("get info successful!", order, ctx)
//		return
//	}
//
//	model.Failed("get info fail", ctx)
//
//}

func GetSingeOrderPage(c *gin.Context) {
	adminId := c.GetUint64("primary_id")
	var page Page
	if c.ShouldBindQuery(&page) != nil {
		model.Failed("bind error", c)
		return
	}
	//fmt.Println(page.PageIndex, page.PageSize)
	//fmt.Println(adminId)
	var orderService service.Order
	if list, ok := orderService.GetSingePage(page.PageSize, page.PageIndex, adminId); ok {
		//fmt.Println(list)
		model.SuccessPage("get list successful!", list, orderService.GetTotal(adminId), c)
		return
	}
	model.Failed("get list error", c)
}
