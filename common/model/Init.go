package model

import (
	"fmt"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type MySql struct {
	DB *gorm.DB
}

var MysqlModel *MySql
var once sync.Once

func NewDb(config map[string]string){
	once.Do(func() {
		MysqlModel = &MySql{}
		//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
		path := strings.Join([]string{config["username"], ":", config["password"], "@tcp(", config["ip"], ":", config["dbPort"], ")/", config["dbName"], "?parseTime=True&loc=Local"}, "")

		//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
		DB, err := gorm.Open("mysql", path)

		if err !=nil {
			fmt.Printf("connnect err: %v", err.Error())
			return
		}

		MysqlModel.DB = DB
	})
}