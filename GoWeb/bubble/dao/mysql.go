package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySql() (err error) {

	dsn := "root:Yu2613309.@tcp(127.0.0.1:3306)/GOTEST?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
	}
	return DB.DB().Ping()
}

func Close() {
	DB.Close()
}
