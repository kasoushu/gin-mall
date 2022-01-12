package service

import "gin_mall/model"

func InsertCategory(c model.Category) bool {
	return true
}
func DeleteCategory(id uint64) bool {
	return true
}

func UpdateCategory(id uint64, c model.Category) bool {
	return true
}
func GetCategoryInfo(id uint64, c *model.Category) bool {
	return true
}
