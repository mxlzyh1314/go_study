package main

/* 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，
并将结果映射到 Book 结构体切片中，确保类型安全。 */

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book  struct {
	ID         	int    	`db:"id"`
	Title       string 	`db:"title"`
	Author 		string 	`db:"author"`
	Price      	int 	`db:"price"`
}

func createTable(db *sqlx.DB) (err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS books(
	id INT AUTO_INCREMENT PRIMARY KEY, 
	title VARCHAR(50), 
	author VARCHAR(50), 
	price INT)`
	_, err = db.Exec(sqlStr)

	return err
}

func insertTable(db *sqlx.DB) { 
	db.Exec("INSERT INTO books VALUES (NULL, 'Go语言趣学指南', '内森杨曼', '69')")
	db.Exec("INSERT INTO books VALUES (NULL, 'Go语言编程', '许式伟', '49')")
	db.Exec("INSERT INTO books VALUES (NULL, 'Go程序设计语言', '艾伦', '79')")
	db.Exec("INSERT INTO books VALUES (NULL, 'Go语言实战', '威廉肯尼迪', '59')")

}

func selectTable(db *sqlx.DB) {
    var books []Book
    err := db.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
    if err != nil {
        fmt.Println("查询失败:", err)
        return
    }
    fmt.Println(books)
}



func main() {
	dsn := "root:123456@tcp(192.168.100.121:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sqlx.Open("mysql", dsn)
	fmt.Println(db, err)
	if err != nil {
		fmt.Println("数据库连接失败") // 打印错误信息
	}

	// createTable(db)
	// insertTable(db)
	selectTable(db)   // 查询价格大于 50 元的书籍

}


