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
}

func (this *User) Validate(username string, password string) bool {
	err := MysqlModel.DB.Table("dumi_xiaoduapp_gateway_env_user").Where("username = ? and password = ?", username, password).Find(this).Error
	return err == nil
}

func (this *User) GetUserById(userId string) bool {
	err := MysqlModel.DB.Table("dumi_xiaoduapp_gateway_env_user").Where("id = ?", userId).Find(this).Error
	return err == nil
}

func (this *User) Update() bool {
	err := MysqlModel.DB.Table("dumi_xiaoduapp_gateway_env_user").Save(this).Error
	return err == nil
}

func (this *User) GetAllUser() ([]User, bool) {
	var userList []User
	err := MysqlModel.DB.Table("dumi_xiaoduapp_gateway_env_user").Find(&userList).Error
	return userList, err == nil
}
