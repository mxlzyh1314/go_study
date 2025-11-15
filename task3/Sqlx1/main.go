package main

/* 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。 */

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Employee  struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     string `db:"salary"`
}

func createTable(db *sqlx.DB) (err error) {
	sqlStr := `CREATE TABLE IF NOT EXISTS employees(
	id INT AUTO_INCREMENT PRIMARY KEY, 
	name VARCHAR(30), 
	department VARCHAR(30), 
	salary VARCHAR(30))`
	_, err = db.Exec(sqlStr)

	return err
}

func insertTable(db *sqlx.DB) { 
	db.Exec("INSERT INTO employees VALUES (NULL, '张三', '技术部', '7000')")
	db.Exec("INSERT INTO employees VALUES (NULL, '李四', '后勤部', '5000')")
	db.Exec("INSERT INTO employees VALUES (NULL, '王五', '人事部', '5500')")
	db.Exec("INSERT INTO employees VALUES (NULL, '赵六', '技术部', '7500')")
	db.Exec("INSERT INTO employees VALUES (NULL, '麻七', '技术部', '7300')")
}

func selectTable(db *sqlx.DB) { 
	var emp []Employee
	db.Select(&emp, "SELECT * FROM employees WHERE department = ?", "技术部")
	fmt.Println(emp)
}

func selectHighestSalaryEmployee(db *sqlx.DB) { 
	var emp Employee
	db.Get(&emp, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	fmt.Println(emp)
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
	// selectTable(db)   // 查找部门为“技术部”的员工信息
	selectHighestSalaryEmployee(db)  // 查找工资最高的员工信
}
