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
		status, created, updated,created_by
		) 
	values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
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
		p.Status, p.Created, p.Updated, p.CreatedBy)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true

}

func DeleteCommodity(id uint64) bool {
	_, err := global.MDB.Exec("delete from products where product_id = ? ", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateCommodity(id uint64, p model.Product) bool {

	//row := global.MDB.QueryRow(`select category_id,
	//   title, description,price,
	//   amount,sales,main_image, delivery,
	//   assurance,name, weight, brand, origin, shelf_life,
	//    net_weight, use_way, packing_way,
	//    storage_condition, detail_image,
	//	status, created, updated from products where id = ? `, id)
	row := global.MDB.QueryRow(`select product_id from products where ?`, id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	p.Updated = time.Now().Format("2006-01-02 15:04:05")

	pre, err := global.MDB.Prepare(`update products set category_id = ? , 
                         title =? ,
                         description=? ,
                         price=? ,
	    				amount=? ,
                         sales=? ,
                         main_image=? ,
                         delivery=? ,
                         assurance=? ,
	    				name=? ,
                         weight=?,
                         brand=? ,
                         origin=? ,
                         shelf_life=? , 
	    net_weight=? , use_way=? , packing_way=? , 
	    storage_condition=? , detail_image=? ,
		status=? , updated=? where product_id = ? `)
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
		p.Status, p.Updated)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GetProductInfo(id uint64, p *model.Product) bool {
	row := global.MDB.QueryRow(`select product_id,category_id,
	  title, description,price,
	  amount,sales,main_image, delivery,
	  assurance,name, weight, brand, origin, shelf_life,
	   net_weight, use_way, packing_way,
	   storage_condition, detail_image,
		status, created, updated,created_by from products where product_id = ? `, id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	err := row.Scan(&p.ProductId, &p.CategoryId, &p.Title, &p.Description, &p.Price,
		&p.Amount, &p.Sales, &p.MainImage, &p.Delivery,
		&p.Assurance, &p.Name, &p.Weight, &p.Brand,
		&p.Origin, &p.ShelfLIfe, &p.NetWeight, &p.UseWay,
		&p.PackingWay, &p.StorageCondition, &p.DetailImage,
		&p.Status, &p.Created, &p.Updated, &p.CreatedBy)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GetTotal(id uint64) int {
	row := global.MDB.QueryRow("select count(*) from products where products.created_by=? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return 0
	}
	var c = 0
	row.Scan(&c)
	return c
}

func GetSinglePageProducts(pageSize, pages int, id uint64) ([]*model.ProductTransfer, bool) {
	var products = make([]*model.ProductTransfer, 0, pageSize)
	rows, err := global.MDB.Query(`select product_id,category_id,categories.name as category_name ,title, description,price,
	  amount,sales,main_image, delivery,
	  assurance,products.name, weight, brand, origin, shelf_life,
	   net_weight, use_way, packing_way,
	   storage_condition, detail_image,
		status, products.created, products.updated,
       created_by from products,categories,admins where (products.created_by=admins.id) and (admins.id=? ) and (categories.id=products.category_id)  limit ? offset ? ;`,
		id, pageSize, (pages-1)*pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	for rows.Next() {
		var p model.ProductTransfer
		err := rows.Scan(&p.ProductId, &p.CategoryId, &p.CategoryName, &p.Title, &p.Description, &p.Price,
			&p.Amount, &p.Sales, &p.MainImage, &p.Delivery,
			&p.Assurance, &p.Name, &p.Weight, &p.Brand,
			&p.Origin, &p.ShelfLIfe, &p.NetWeight, &p.UseWay,
			&p.PackingWay, &p.StorageCondition, &p.DetailImage,
			&p.Status, &p.Created, &p.Updated, &p.CreatedBy)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		//fmt.Println(p)
		products = append(products, &p)
	}
	return products, true
}
