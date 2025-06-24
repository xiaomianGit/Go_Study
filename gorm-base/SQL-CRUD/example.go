package SQL_CRUD

import (
	"fmt"
	"gorm-base/DO"
	"gorm.io/gorm"
)

/*
*
*题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func Run(db *gorm.DB) {
	createTable(db)
	insertRecord(db)
	selectRecord(db)
	updateRecord(db)
	deleteRecord(db)
}

func deleteRecord(db *gorm.DB) {
	db.Where("Age < ?", 15).Delete(&DO.Students{})
}

func updateRecord(db *gorm.DB) {
	db.Model(&DO.Students{}).Where("Name = ?", "张三").Update("Grade", "四年级")
}

func selectRecord(db *gorm.DB) {
	var students []DO.Students
	result := db.Where("age > ?", 20).Find(&students)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, student := range students {
		fmt.Printf("student:%+v \n", student)
	}
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func insertRecord(db *gorm.DB) {
	db.Create(&DO.Students{Name: "张三", Age: 20, Grade: "三年级"})
	db.Create(&DO.Students{Name: "李四", Age: 25, Grade: "三年级"})
	db.Create(&DO.Students{Name: "王五", Age: 12, Grade: "三年级"})
}
func createTable(db *gorm.DB) {
	db.AutoMigrate(DO.Students{})
}
