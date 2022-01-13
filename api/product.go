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
	if product.Name == "" || product.Amount == 0 || product.Price == 0 {
		model.Failed("params error", ctx)
		return
	}

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
