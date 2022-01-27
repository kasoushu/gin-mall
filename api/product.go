package api

import (
	"fmt"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreteProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}

	fmt.Println(fmt.Sprintf("name:%s,des:%s,price:%f,amout%f", product.Name, product.Description, product.Price, product.Amount))
	if product.Name == "" || product.Amount == 0 || product.Price == 0 {
		model.Failed("params can not be empty", ctx)
		return
	}
	id := ctx.MustGet("primary_id")
	product.CreatedBy = id.(uint64)
	if service.CreateCommodity(product) {
		model.Success("create successful!", "", ctx)
		return
	}
	model.Failed("create fail", ctx)
}

func UpdateProduct(ctx *gin.Context) {
	var product model.Product

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if err := ctx.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if service.UpdateCommodity(uint64(id), product) {
		model.Success("update successful!", "", ctx)
		return
	}
	model.Failed("update fail", ctx)
}
func DeleteProduct(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if service.DeleteCommodity(id) {
		model.Success("delete successful!", "", ctx)
		return
	}
	model.Failed("delete fail", ctx)
}

func GetSingleProductInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	var product model.Product
	if service.GetProductInfo(id, &product) {
		//p, err := json.Marshal(product)
		if err != nil {
			fmt.Println(err)
			model.Failed("convert json error", ctx)
		}
		model.Success("get info successful!", product, ctx)
		return
	}

	model.Failed("get info fail", ctx)

}

type Page struct {
	PageSize  int ` form:"page_size"`
	PageIndex int `form:"page_index"`
}

func GetSingePage(c *gin.Context) {
	adminId := c.GetUint64("primary_id")
	var page Page
	if c.ShouldBindQuery(&page) != nil {
		model.Failed("bind error", c)
		return
	}

	fmt.Println(page.PageIndex, page.PageSize)
	fmt.Println(adminId)

	if list, ok := service.GetSinglePageProducts(page.PageSize, page.PageIndex, adminId); ok {
		fmt.Println(list)
		model.SuccessPage("get list successful!", list, service.GetTotal(adminId), c)
		return
	}
	model.Failed("get list error", c)
}

func GetStatisticByStatus(ctx *gin.Context) {

	productStatistic, err := service.GetProductByStatus()
	if err != nil {
		fmt.Println(err)
		model.Failed("get statistic failed", ctx)
		return
	}

	for _, p := range productStatistic.List {
		productStatistic.Total += p.Total
	}

	model.Success("get status astatistic successful!", productStatistic, ctx)
}
