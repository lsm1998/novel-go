package config

import (
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"utils"
)

var Config utils.Config

var (
	DB          *gorm.DB
	MysqlDB     *sql.DB
	RedisClient *redis.Pool
)

func Init() {
	mysqlC := Config.Mysql
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, mysqlC.User, mysqlC.Password, mysqlC.Url, mysqlC.Db))
	if err != nil {
		panic(err)
	}
	DB = db
	MysqlDB = db.DB()
	DB.LogMode(true)

	redisC := Config.Redis
	RedisClient = &redis.Pool{
		MaxIdle:     redisC.MaxIdle,
		MaxActive:   redisC.MaxActive,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf(`%s:%d`, redisC.Adder, redisC.Port))
			if err != nil {
				return nil, err
			}
			c.Do("auth", redisC.Auth)
			c.Do("SELECT", redisC.Db)
			return c, nil
		},
	}
}
