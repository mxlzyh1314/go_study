package database

import (
	"fmt"
	. "task4/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	fmt.Println("数据库初始化...")
	dsn := "root:123456@tcp(127.0.0.1:3306)/task4?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	log.Println("数据库连接成功")
}
