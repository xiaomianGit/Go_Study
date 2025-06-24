package SQLX_extendQuery

import (
	"fmt"
	"gorm-base/DO"
	"gorm.io/gorm"
)

/*
*题目1：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func Run(db *gorm.DB) {
	createTable(db)
	insertRecord(db)
	selectRecord(db)
	selectRecord2(db)
}

func selectRecord2(db *gorm.DB) {
	var employees []DO.Employees
	subQuery := db.Model(&DO.Employees{}).Select("max(salary)")
	result := db.Where("salary = (?)", subQuery).Find(&employees)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, ele := range employees {
		fmt.Printf("工资最高 employees: %+v\n", ele)
	}
}

func selectRecord(db *gorm.DB) {
	var employees []DO.Employees
	result := db.Where("department = ?", "技术部").Find(&employees)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, ele := range employees {
		fmt.Printf("技术部 employees:%+v \n", ele)
	}
}

func insertRecord(db *gorm.DB) {
	db.Create(&DO.Employees{Name: "张三", Department: "技术部", Salary: 20000.01})
	db.Create(&DO.Employees{Name: "李四", Department: "技术部", Salary: 10000.01})
	db.Create(&DO.Employees{Name: "王五", Department: "技术部", Salary: 30000.01})
}
func createTable(db *gorm.DB) {
	db.AutoMigrate(DO.Employees{})
}
