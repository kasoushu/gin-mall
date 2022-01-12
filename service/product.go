package service

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"time"
)

func CreateCommodity(p model.Product) bool {
	p.Created = time.Now().Format("2006-01-02 15:04:05")
	p.Updated = time.Now().Format("2006-01-02 15:04:05")
	pre, err := global.MDB.Prepare(`
	insert into  products (category_id, title, description, price,
	    amount,sales, main_image, delivery, assurance,
	    name, weight, brand, origin, shelf_life, 
	    net_weight, use_way, packing_way, 
	    storage_condition, detail_image,
		status, created, updated
		) 
	values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(
		p.CategoryId, p.Title, p.Description, p.Price,
		p.Amount, p.Sales, p.MainImage, p.Delivery,
		p.Assurance, p.Name, p.Weight, p.Brand,
		p.Origin, p.ShelfLIfe, p.NetWeight, p.UseWay,
		p.PackingWay, p.StorageCondition, p.DetailImage,
		p.Status, p.Created, p.Updated)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
