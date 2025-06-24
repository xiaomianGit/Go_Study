package main

import (
	"fmt"
	High_modelDefine "gorm-base/High-modelDefine"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:11!!qqQQ@tcp(127.0.0.1:3306)/gorm-base?charset=utf8mb4&parseTime=True&loc=Local&sql_mode=ALLOW_INVALID_DATES"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	fmt.Printf("链接成功=============%v \n", db)

	//lession01.Run(db)
	//SQL_CRUD.Run(db)
	//SQL_translation.Run(db)
	//SQLX_extendQuery.Run(db)
	//SQLX_typeMapping.Run(db)
	High_modelDefine.Run(db)
}
