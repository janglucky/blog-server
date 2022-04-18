package model

import (
	"time"
)

type Article struct {
	Id      string    `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Ctime   time.Time `json:"ctime"`
	Utime   time.Time `json:"utime"`
	Status  int      `json:"status"`
}

type Tag struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ctime       time.Time `json:"ctime"`
	Utime       time.Time `json:"utime"`
}

type ArticleTagRel struct {
	ArticleId string `json:"article_id"`
	TagId string `json:"tag_id"`
}

func (this *Article) Create() error {
	err := MysqlModel.DB.Table("tensir_blog_article").Create(this).Error
	return err
}

func (this *Tag) Create() error {
	err := MysqlModel.DB.Table("tensir_blog_tag").Create(this).Error
	return err
}
func (this *ArticleTagRel) Create() error {
	err := MysqlModel.DB.Table("tensir_blog_article_tag").Create(this).Error
	return err
}


func (this *Article) GetUserById(userId string) bool {
	err := MysqlModel.DB.Table("tensir_blog_user").Where("id = ?", userId).Find(this).Error
	return err == nil
}

func (this *Article) Update() bool {
	err := MysqlModel.DB.Table("tensir_blog_user").Save(this).Error
	return err == nil
}

func (this *Article) GetAllUser() ([]User, bool) {
	var userList []User
	err := MysqlModel.DB.Table("tensir_blog_user").Find(&userList).Error
	return userList, err == nil
}
