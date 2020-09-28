package config

import (
	"database/sql"
	"fmt"
	"github.com/Shopify/sarama"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"utils"
)

var Config utils.Config
var KafkaProducer sarama.SyncProducer

var DB *gorm.DB
var MysqlDB *sql.DB

var SnowFlake = new(utils.SnowFlake)

func Init() {
	// ---------kafka--------
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	var err error
	if KafkaProducer, err = sarama.NewSyncProducer(Config.Im.Adders, config); err != nil {
		panic(err)
	}

	// ---------Mysql--------
	mysql := Config.Mysql
	db, err := gorm.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local`, mysql.User, mysql.Password, mysql.Url, mysql.Db))
	if err != nil {
		panic(err)
	}
	DB = db
	MysqlDB = db.DB()
	DB.LogMode(true)
}
