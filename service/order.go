package service

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"time"
)

func GetOrderInfo(id uint64, order *model.Order) bool {
	row := global.MDB.QueryRow(`select id, product_item, total_price,
			address_id, user_id, nick_name, created,
       updated, product_id,status from orders where orders.id = ? `, id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	err := row.Scan(&order.Id, &order.ProductItem, &order.TotalPrice,
		&order.AddressId, &order.UserId,
		&order.NickName, &order.Created, &order.Updated, &order.ProductId, &order.Status)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateOrder(id uint64, order model.Order) bool {
	order.Updated = time.Now().Format("2006-01-02 15:04:05")
	row := global.MDB.QueryRow("select orders.id from orders where orders.id=? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	pre, err := global.MDB.Prepare("update orders set status=? ,orders.updated = ? where orders.id=? ")
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(order.Status, order.Updated, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func DeleteOrder(id uint64) bool {
	_, err := global.MDB.Exec("delete from orders where orders.id=? ", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func InsertOrder(order model.Order) bool {
	order.Created = time.Now().Format("2006-01-02 15:04:05")
	order.Updated = time.Now().Format("2006-01-02 15:04:05")
	pre, err := global.MDB.Prepare(`insert into orders (product_item, total_price,
			address_id, user_id, nick_name, created,
       updated, product_id,status ) values (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(order.ProductItem, order.TotalPrice,
		order.AddressId, order.UserId, order.NickName,
		order.Created, order.Updated, order.ProductId, order.Status)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
