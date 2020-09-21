package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"user/config"
)

var DB *gorm.DB

func init() {
	mysql := config.Config.Mysql
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@/%s?charset=utf8&parseTime=True&loc=Local`, mysql.User, mysql.Password, mysql.Db))
	if err != nil {
		panic(err)
	}
	DB = db
}
