package High_modelDefine

import (
	"fmt"
	"gorm-base/DO"
	"gorm.io/gorm"
)

/**
题目1：模型定义
	假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
	要求 ：
	使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
	编写Go代码，使用Gorm创建这些模型对应的数据库表。
题目2：关联查询
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
题目3：钩子函数
	继续使用博客系统的模型。
	要求 ：
	为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

func Run(db *gorm.DB) {
	createTable(db)
	//insertRecord(db)
	selectRecord(db)
	selectRecord2(db)
	deleteRecord(db)
}

func deleteRecord(db *gorm.DB) {
	var comments []DO.Comment
	db.Where("post_id = ?", 1).Find(&comments)
	for i, _ := range comments {
		// 通过索引访问原始切片元素的指针
		if i == 0 {
			comment := &comments[i]
			db.Delete(comment) // 传递原始对象的指针，触发钩子
		}
	}
}
func createTable(db *gorm.DB) {
	db.AutoMigrate(DO.User{}, DO.Post{}, DO.Comment{})
}
func insertRecord(db *gorm.DB) {
	db.Create(&DO.User{Username: "张三"})
	db.Create(&DO.Post{Title: "go从入门到放弃", UserId: 1})
	db.Create(&DO.Post{Title: "java从入门到放弃", UserId: 1})
	db.Create(&DO.Comment{Content: "有点东西", UserId: 1, PostId: 1})
	db.Create(&DO.Comment{Content: "但不多", UserId: 1, PostId: 1})
}

/**
题目2：关联查询
	基于上述博客系统的模型定义。
	要求 ：
	编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func selectRecord(db *gorm.DB) {
	var posts []DO.Post
	db.Where("user_id = ?", 1).Model(DO.Post{}).Find(&posts)

	for _, post := range posts {
		var comments []DO.Comment
		db.Where("post_id = ?", post.Id).Model(DO.Comment{}).Find(&comments)
		fmt.Printf("post %v comments %v\n", post, comments)
	}
}
func selectRecord2(db *gorm.DB) {
	var posts []DO.Post
	db.Where("user_id = ?", 1).Order("Comment_num desc").Model(DO.Post{}).First(&posts)

	fmt.Printf("post %v comments %v\n", posts[0])
}
