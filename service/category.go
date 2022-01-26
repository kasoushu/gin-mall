package service

import (
	"database/sql"
	"fmt"
	"gin_mall/global"
	"gin_mall/model"
	"time"
)

func InsertCategory(c model.Category) bool {
	row := global.MDB.QueryRow("select id from categories where name=? ", c.Name)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	flag := -1
	if row.Scan(&flag); flag != -1 {
		fmt.Println("already existed")
		return false
	}
	//set time
	c.Created = time.Now().Format("2006-01-02 15:04:05")
	c.Updated = time.Now().Format("2006-01-02 15:04:05")
	pre, err := global.MDB.Prepare(`insert into categories (name, parent_id,
                         created, updated ) values (?,?,?,?,?,?) `)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = pre.Exec(c.Name, c.ParentId, c.Created, c.Updated)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func DeleteCategory(id uint64) bool {
	_, err := global.MDB.Exec("delete from categories where categories.id=? ", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateCategory(id uint64, c model.Category) bool {
	row := global.MDB.QueryRow("select id from categories where id=? ", id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	flag := -1
	if row.Scan(&flag); flag == -1 {
		fmt.Println("already not existed")
		return false
	}
	c.Updated = time.Now().Format("2006-01-02 15:04:05")
	_, err := global.MDB.Exec(`update categories set name =?,
					parent_id=?  ,updated=? where categories.id=? `, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GetCategoryInfo(id uint64, c *model.Category) bool {

	pre, err := global.MDB.Prepare(`select id, name, parent_id, created, updated from categories where categories.id=? `)
	if err != nil {
		fmt.Println(err)
		return false
	}
	row := pre.QueryRow(id)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return false
	}
	err = row.Scan(&c.Id, &c.Name, &c.ParentId, &c.Created, &c.Updated)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
func GetParentCategories(pid uint64) *sql.Rows {
	rows, err := global.MDB.Query("select * from categories where parent_id=? ", pid)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows
}
