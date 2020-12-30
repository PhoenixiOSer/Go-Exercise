package main

import (
	"gin_webapp/dao"
	"gin_webapp/models"
	"gin_webapp/routers"
)

func main() {

	// 创建数据库 sql： CREATE TABLE bubble

	// 连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	if err := r.Run(":9090"); err != nil {

	}
}
