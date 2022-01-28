package service

import (
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"time"
)

func InsertAddress(address model.Address) bool {

	address.Created = time.Now().Format("2006-01-02 15:04:05")
	address.Updated = time.Now().Format("2006-01-02 15:04:05")
	pre, err := global.MDB.Prepare(`insert into address(name, user_id,
                    postal_code, created, updated, province,
                    city, district, detail_address,is_default) values (?,?,?,?,?,?,?,?,?,?,? ) `)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(address.Name, address.UserId, address.PostalCode, address.Created,
		address.Updated, address.Province, address.City, address.District, address.DetailAddress, address.IsDefault)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func UpdateAddress(id uint64, address model.Address) bool {

	address.Updated = time.Now().Format("2006-01-02 15:04:05")

	row := global.MDB.QueryRow("select address.id from address where address.id=? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	pre, err := global.MDB.Prepare(`update address set name=? ,
                    postal_code=? , updated=? , province=? ,
                    city=? , district=? , detail_address=? ,is_default=? where address.id=? `)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(address.Name, address.PostalCode,
		address.Updated, address.Province, address.City, address.District,
		address.DetailAddress, address.IsDefault)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func DeleteAddress(address_id uint64) bool {
	_, err := global.MDB.Exec("delete  from address where address.id=? ", address_id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GetAddressInfo(address_id uint64) (*model.AddressTransfer, bool) {
	var address model.AddressTransfer
	pre, err := global.MDB.Prepare(`select address.id,address.name, users.name,users.phone,postal_code, province,
                    city, district, detail_address,is_default from address,users where (address.user_id=users.id) and (address.id=? ) `)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	row := pre.QueryRow(address_id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return nil, false
	}
	err = row.Scan(&address.Id, &address.Name, &address.User.Name, &address.User.Phone, &address.PostalCode,
		&address.Province, &address.City, &address.District, &address.DetailAddress, &address.IsDefault)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}
	return &address, true
}
