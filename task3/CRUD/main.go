package main
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Students struct {
	Id      int        `gorm:"primaryKey"`
	Name 	string
	Agent   int
	Grade   string
}

func addRecord(db *gorm.DB) {
	db.Create(&Students{Name: "张三", Agent: 20, Grade: "三年级"})
}

func selectRecord(db *gorm.DB) {
	var students []Students
	db.Where("agent > ?",18).Find(&students)
	fmt.Println(students)
}

func updateRecord(db *gorm.DB) { 
	db.Model(&Students{}).Where("name = ?", "张三").Update("grade", "四年级")
}

func deleteRecord(db *gorm.DB) { 
	db.Where("agent < ?", 15).Delete(&Students{})
}
func main () { 
	dsn := "root:123456@tcp(192.168.100.121:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	if err != nil {
		fmt.Println("数据库连接失败") // 打印错误信息
	}

	// 迁移表
	db.AutoMigrate(&Students{})

	// addRecord(db)
	// selectRecord(db)
	// updateRecord(db)
	deleteRecord(db)
}