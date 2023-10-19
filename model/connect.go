package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBUser = "root"
	DBPWD  = "TC123"
	DBIP   = "127.0.0.1:3306"
	DBName = "ccy"
)

var DB *gorm.DB
var err error

func init() {
	dsn := DBUser + ":" + DBPWD + "@tcp(" + DBIP + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&B_list{})
}
