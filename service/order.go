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

	row := global.MDB.QueryRow("select status from orders where orders.id= ? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	s := ""
	if err := row.Scan(&s); err != nil {
		fmt.Println(err)
		return false
	}
	if s == "0" {
		fmt.Println("error order is not done")
		return false
	}
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

	rows, err := global.MDB.Query(`select orders.id,products.name,address_id,user_id,
       total_price,orders.status,orders.created,orders.updated,orders.amount from  orders,products where (products.product_id=orders.product_id) and (admin_id=?) order by orders.id desc limit ? offset ? `, id, pageSize, (pages-1)*pageSize)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	for rows.Next() {
		var p model.OrderTransfer
		var (
			uId uint64
			aId uint64
		)
		err := rows.Scan(&p.Id, &p.ProductName, &aId, &uId, &p.TotalPrice, &p.Status,
			&p.Created, &p.Updated, &p.Amount)
		if err != nil {
			fmt.Println(err)
			return nil, false
		}

		conn, err := global.MDB.Begin()
		if err != nil {
			fmt.Println(err)
			return nil, false
		}

		row := conn.QueryRow("select name,phone from users where id=? ", uId)
		if row.Err() != nil {
			fmt.Println(row.Err())
			return nil, false
		}
		if err := row.Scan(&p.UserInfo.Name, &p.UserInfo.Phone); err != nil {
			fmt.Println(err)
			return nil, false
		}
		row = conn.QueryRow(`select address.id,address.name, users.name,users.phone,postal_code, province,
                    city, district, detail_address,is_default from address,users where (address.user_id=users.id) and (address.id=? )  `, aId)
		if row.Err() != nil {
			fmt.Println(row.Err())
			return nil, false
		}
		var address model.AddressTransfer
		if err := row.Scan(&address.Id, &address.Name, &address.User.Name, &address.User.Phone, &address.PostalCode, &address.Province, &address.City, &address.District, &address.DetailAddress, &address.IsDefault); err != nil {
			fmt.Println(err)
			return nil, false
		}
		p.Address = address

		err = conn.Commit()
		//fmt.Println(p)
		orders = append(orders, p)
	}
	if err != nil {
		fmt.Println(err)
		return nil, false
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

func (o *Order) TenDaysOrderCount(id uint64) ([]model.OrderCount, error) {
	t := time.Now().Add(-24 * 10 * time.Hour)
	list := make([]model.OrderCount, 0)
	con, err := global.MDB.Begin()
	if err != nil {
		return nil, err
	}
	for i := 0; i < 10; i++ {
		t = t.Add(24 * time.Hour)
		y, m, d := t.Year(), int(t.Month()), t.Day()
		//fmt.Println(y, " ", m, " ", d)
		var oC model.OrderCount
		row := con.QueryRow("select count(*) from orders where admin_id=?  and  year(created)= ? and  month(created)=? and  day(created)=?  ", id, y, m, d)
		if row.Err() != nil {
			return nil, row.Err()
		}
		if err := row.Scan(&oC.Cnt); err != nil {
			return nil, err
		}
		oC.Day = t.Format("01-02")
		list = append(list, oC)
	}
	if err := con.Commit(); err != nil {
		return nil, err
	}
	return list, nil
}

func (o *Order) OrderStatistic(id uint64) (*model.OrderStatistic, error) {
	var oSta model.OrderStatistic

	con, err := global.MDB.Begin()
	if err != nil {
		return nil, err
	}
	row := con.QueryRow("select count(*) from orders where admin_id=? ", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.OrderTotal); err != nil {
		return nil, err
	}
	t := time.Now()
	y, m, d := t.Year(), int(t.Month()), t.Day()
	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=? ", id, y, m)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.MonthTotal); err != nil {
		return nil, err
	}
	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=? and day(created)=?  ", id, y, m, d)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.DayTotal); err != nil {
		return nil, err
	}

	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=? and day(created)=?  and status=0 ", id, y, m, d)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.DayNotDone); err != nil {
		return nil, err
	}

	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=? and day(created)=?  and status=1 ", id, y, m, d)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.DayDone); err != nil {
		return nil, err
	}

	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=?  and status=1 ", id, y, m)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.MonthDone); err != nil {
		return nil, err
	}

	row = con.QueryRow("select count(*) from orders where admin_id=? and year(created)=? and month(created)=?  and status=0 ", id, y, m)

	if row.Err() != nil {
		return nil, row.Err()
	}
	if err := row.Scan(&oSta.MonthNotDone); err != nil {
		return nil, err
	}
	if err := con.Commit(); err != nil {
		return nil, err
	}
	return &oSta, nil
}
