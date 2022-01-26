package api

import (
	"fmt"
	"gin_mall/model"
	"gin_mall/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreteCategory(ctx *gin.Context) {
	var category model.Category
	if err := ctx.ShouldBindJSON(&category); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if category.Name == "" {
		model.Failed("params error", ctx)
		return
	}

	if service.InsertCategory(category) {
		model.Success("create successful!", "", ctx)
		return
	}
	model.Failed("create fail", ctx)
}

func UpdateCategory(ctx *gin.Context) {
	var category model.Category

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if err := ctx.ShouldBindJSON(&category); err != nil {
		fmt.Println(err)
		model.Failed("params error", ctx)
		return
	}
	if service.UpdateCategory(id, category) {
		model.Success("update successful!", "", ctx)
		return
	}
	model.Failed("update fail", ctx)
}
func DeleteCategory(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	if service.DeleteCategory(id) {
		model.Success("delete successful!", "", ctx)
		return
	}
	model.Failed("delete fail", ctx)
}

func GetCategoryInfo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		model.Failed("parse params error", ctx)
		return
	}
	var category model.Category
	if service.GetCategoryInfo(id, &category) {
		if err != nil {
			fmt.Println(err)
			model.Failed("convert json error", ctx)
		}
		model.Success("get info successful!", category, ctx)
		return
	}
	model.Failed("get info fail", ctx)

}

func GetCategoriesParent(ctx *gin.Context) {
	var list = make([]model.Category, 0)
	pid := ctx.Param("pid")
	id, err := strconv.ParseUint(pid, 10, 64)
	if err != nil {
		model.Failed("parse pid error", ctx)
		return
	}
	rows := service.GetParentCategories(id)
	for rows.Next() {
		var c model.Category
		rows.Scan(&c.Id, &c.Name, &c.ParentId, &c.Created, &c.Updated)
		fmt.Println(c)
		list = append(list, c)
	}
	model.Success("get cateogries parents successgul! ", model.CategoryParents{
		Total: len(list),
		List:  list,
	}, ctx)
}
