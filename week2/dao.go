package db

import (
	"myerror"
)

type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"`
}

func (dao *userDAO) GetUser(name string) (*User, error) {
	var user User
	db := db.Where(&User{Name: name}).First(&user)
	if db.RecordNotFound() {
		return nil, myerror.NewErrNotFound("User", name)
	}
	return &user, db.Error
}
