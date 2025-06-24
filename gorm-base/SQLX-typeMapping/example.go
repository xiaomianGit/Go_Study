package SQLX_typeMapping

import (
	"fmt"
	"gorm-base/DO"
	"gorm.io/gorm"
)

/*
*假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func Run(db *gorm.DB) {
	createTable(db)
	insertRecord(db)
	selectRecord(db)
}
func selectRecord(db *gorm.DB) {
	var books []DO.Books
	result := db.Where("Price > ?", 50).Find(&books)
	if result.Error != nil {
		panic(result.Error)
	}
	for _, ele := range books {
		fmt.Printf("价格大于50的books:%+v \n", ele)
	}
}

func insertRecord(db *gorm.DB) {
	db.Create(&DO.Books{Title: "go 从入门到放弃", Author: "技术部", Price: 100})
	db.Create(&DO.Books{Title: "java 从入门到放弃", Author: "技术部", Price: 60})
	db.Create(&DO.Books{Title: "python 从入门到放弃", Author: "技术部", Price: 40})
}
func createTable(db *gorm.DB) {
	db.AutoMigrate(DO.Books{})
}
