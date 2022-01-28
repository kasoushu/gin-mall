package service

import (
	"gin_mall/common"
	"gin_mall/global"
	"gin_mall/model"
)

type User struct {
}

func IsExist(w model.WebLoginUser, u *model.User) bool {
	row, err := common.QueryUser(w)
	if err != nil {
		return false
	}
	if row != nil {
		err = row.Scan(&u.Id, &u.Name, &u.Phone)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func (user *User) UserInfo(id uint64) (*model.UserInfo, error) {
	var u model.UserInfo
	row := global.MDB.QueryRow("select name,phone from users where users.id=? ", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(&u.Name, &u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
