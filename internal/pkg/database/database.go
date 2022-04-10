package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var dsn string = "root:root@tcp(192.168.31.57:3306)/home_env?charset=utf8mb4&parseTime=True&loc=Local"

func InitDb(m interface{}) *gorm.DB {
	if db == nil {
		tmpDB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		})
		db = tmpDB
	}
	db.AutoMigrate(m)

	return db
}
