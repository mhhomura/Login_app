package models

import (
	"github.com/jinzhu/gorm"
)

const (
	USERS = "users"
)

func GetAll(table string) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Order("id asc").Find(&[]User{}).Value
	}
	return nil
}

func GetById(table string, id uint64) interface{} {
	db := Connect()
	defer db.Close()
	switch table {
	case USERS:
		return db.Where("id = ?", id).Find(&User{}).Value
	}
	return nil
}

func Delete(table string, id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	var rs *gorm.DB
	switch table {
	case USERS:
		rs = db.Where("id = ?", id).Delete(&User{})
	}
	return rs.RowsAffected, rs.Error
}
