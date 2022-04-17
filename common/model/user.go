package model

import (
	"time"
)

type User struct {
	Id       string    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Ctime    time.Time `json:"ctime"`
	Utime    time.Time `json:"utime"`
	Avatar   string    `json:"avatar"`
	email    string    `json:"email"`
}

func (this *User) Validate() error {
	err := MysqlModel.DB.Table("tensir_blog_user").Where("username = ? and password = ?", this.Username, this.Password).Find(this).Error
	return err
}

func (this *User) GetUserById(userId string) bool {
	err := MysqlModel.DB.Table("tensir_blog_user").Where("id = ?", userId).Find(this).Error
	return err == nil
}

func (this *User) Update() bool {
	err := MysqlModel.DB.Table("tensir_blog_user").Save(this).Error
	return err == nil
}

func (this *User) GetAllUser() ([]User, bool) {
	var userList []User
	err := MysqlModel.DB.Table("tensir_blog_user").Find(&userList).Error
	return userList, err == nil
}
