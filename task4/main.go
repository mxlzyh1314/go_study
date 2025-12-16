package main

import (
	"task4/config"
	"task4/database"
	"task4/routers"
	"task4/utils"
)

func main() {
	// 加载配置
	config.LoadConfig()
	// 初始化日志
	utils.InitLogger()
	// 初始化数据库
	database.InitMysql()
	// 启动服务
	r := routers.SetupRouter()
	r.Run(":8080")
}
