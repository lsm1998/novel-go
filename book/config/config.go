package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"utils"
)

var Config utils.Config

var DB *gorm.DB
var MysqlDB *sql.DB

func Init() {
	mysql := Config.Mysql
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, mysql.User, mysql.Password, mysql.Url, mysql.Db))
	if err != nil {
		panic(err)
	}
	DB = db
	MysqlDB = db.DB()
}
