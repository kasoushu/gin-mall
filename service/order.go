package service

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"time"
)

type Order struct{}

//func GetOrderInfo(id uint64, order *model.Order) bool {
//	row := global.MDB.QueryRow(`select id, product_item, total_price,
//			address_id, user_id, nick_name, created,
//       updated, product_id,status from orders where orders.id = ? `, id)
//	if row.Err() != nil {
//		fmt.Println(row.Err())
//		return false
//	}
//	err := row.Scan(&order.Id, &order.ProductItem, &order.TotalPrice,
//		&order.AddressId, &order.UserId,
//		&order.NickName, &order.Created, &order.Updated, &order.ProductId, &order.Status)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	return true
//}

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
	pre, err := global.MDB.Prepare(`insert into orders (total_price,
			address_id, user_id,  created,
       updated, product_id,status,admin_id,amount ) values (?,?,?,?,?,?,?,?,?)`)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(order.TotalPrice,
		order.AddressId, order.UserId,
		order.Created, order.Updated, order.ProductId, order.Status, order.AdminId, order.Amount)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (o *Order) GetSingePage(pageSize, pages int, id uint64) ([]model.OrderTransfer, bool) {
	var orders = make([]model.OrderTransfer, 0, pageSize)

	rows, err := global.MDB.Query(`select orders.id,product_id,address_id,user_id,
       total_price,status,created,updated,amount from  orders where admin_id=? limit ? offset ? `, id, pageSize, (pages-1)*pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	for rows.Next() {
		var p model.OrderTransfer
		var (
			pId uint64
			uId uint64
			aId uint64
		)
		err := rows.Scan(&p.Id, &pId, &aId, &uId, &p.TotalPrice, &p.Status,
			&p.Created, &p.Updated, &p.Amount)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		var userService User
		u, err := userService.UserInfo(uId)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}
		p.UserInfo = *u
		if s, ok := GetProductName(pId); ok {
			p.ProductName = s
		} else {
			return nil, false
		}
		if add, ok := GetAddressInfo(aId); ok {
			p.Address = *add
		} else {
			return nil, false
		}
		//fmt.Println(p)
		orders = append(orders, p)
	}

	return orders, true
}

func (o *Order) GetTotal(id uint64) int {
	row := global.MDB.QueryRow("select count(*) from orders where orders.admin_id =? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return 0
	}
	var c = 0
	row.Scan(&c)
	return c
}
