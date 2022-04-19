package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Article struct {
	Id       string     `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Ctime    *time.Time `json:"ctime"`
	Utime    *time.Time `json:"utime"`
	Status   int        `json:"status"`
	AuthorId string     `json:"author_id"`
}

type Tag struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Ctime       *time.Time `json:"ctime"`
	Utime       *time.Time `json:"utime"`
}

type ArticleTagRel struct {
	ArticleId string `json:"article_id"`
	TagId     string `json:"tag_id"`
}

func (this *Article) SetAuthorId(authorId string)  {
	if this != nil {
		this.AuthorId = authorId
	}
}
func (this *Article) SetCtime()  {
	if this == nil {
		return
	}
	if this.Ctime == nil {
		this.Ctime = new(time.Time)
	}
	*this.Ctime = time.Now()
}

func (this *Article) Create() error {
	err := MysqlModel.DB.Table("tensir_blog_article").Create(this).Error
	return err
}

func (this *Article) String() string {
	b, err := json.Marshal(*this)
	if err != nil {
		return fmt.Sprintf("%+v", *this)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *this)
	}
	return out.String()
}

func (this *Tag) Create() error {
	err := MysqlModel.DB.Table("tensir_blog_tag").Create(this).Error
	return err
}

// SearchByKeyword ...按关键字搜索标签
func (this *Tag) SearchByKeyword(keyword string) ([]Tag, error) {
	// todo 使用索引
	var tags []Tag
	err := MysqlModel.DB.Table("tensir_blog_tag").Where("name like ?", keyword+"%").Find(&tags).Error
	return tags,err
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
