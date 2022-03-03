package api

import (
	"encoding/json"
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CreteOrder(ctx *gin.Context) {
	adminId := ctx.GetUint64("primary_id")
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

		adminKey := global.AdminOrderListPrefix + strconv.FormatUint(adminId, 10)
		_, err := global.RDB.Del(adminKey).Result()
		if err != nil {
			model.Failed("create fail", ctx)
		}
		model.Success("create successful!", "", ctx)
		return
	}
	model.Failed("create fail", ctx)
}

func UpdateOrder(ctx *gin.Context) {
	var order model.Order
	adminId := ctx.GetUint64("primary_id")
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
		adminKey := global.AdminOrderListPrefix + strconv.FormatUint(adminId, 10)
		_, err := global.RDB.Del(adminKey).Result()
		if err != nil {
			model.Failed("update fail", ctx)
		}
		model.Success("update successful!", "", ctx)
		return
	}
	model.Failed("update fail", ctx)
}
func DeleteOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	adminId := ctx.GetUint64("primary_id")
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if service.DeleteOrder(id) {
		adminKey := global.AdminOrderListPrefix + strconv.FormatUint(adminId, 10)
		_, err := global.RDB.Del(adminKey).Result()
		if err != nil {
			model.Failed("update fail", ctx)
		}
		model.Success("delete successful!", "true", ctx)
		return
	}
	model.Failed("delete fail order is not done or network error", ctx)
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
	adminKey := global.AdminOrderListPrefix + strconv.FormatUint(adminId, 10)
	var orderService service.Order
	var list []model.OrderTransfer
	var page Page
	if c.ShouldBindQuery(&page) != nil {
		model.Failed("bind error", c)
		return
	}
	//fmt.Println(page.PageIndex, page.PageSize)
	//fmt.Println(adminId)
	if global.RDB.Exists(adminKey).Val() == 1 {
		ls, err := global.RDB.Get(adminKey).Result()
		if err != nil {
			fmt.Println(err)
			model.Failed("get list error", c)
			return
		}
		err = json.Unmarshal([]byte(ls), &list)
		if err != nil {
			fmt.Println(err)
			model.Failed("get list error", c)
			return
		}
		model.SuccessPage("get list successful!", list, orderService.GetTotal(adminId), c)
		return
	}
	if list, ok := orderService.GetSingePage(page.PageSize, page.PageIndex, adminId); ok {
		//fmt.Println(list)
		lstr, _ := json.Marshal(list)
		//fmt.Println(string(lstr))
		global.RDB.Set(adminKey, string(lstr), time.Hour)
		model.SuccessPage("get list successful!", list, orderService.GetTotal(adminId), c)
		return
	}
	model.Failed("get list error", c)
}

func GetTenDaysOrderCount(ctx *gin.Context) {
	adminId := ctx.GetUint64("primary_id")
	var orderService service.Order

	m, err := orderService.TenDaysOrderCount(adminId)
	if err != nil {
		fmt.Println(err)
		model.Failed("get data error ", ctx)
		return
	}

	model.Success("get data successful!", m, ctx)
}

func GetOrderStatistic(ctx *gin.Context) {
	adminId := ctx.GetUint64("primary_id")
	var orderService service.Order
	st, err := orderService.OrderStatistic(adminId)
	if err != nil {
		fmt.Println(err)
		model.Failed("get message error", ctx)
		return
	}
	model.Success("get order statistic successful! ", st, ctx)
}
