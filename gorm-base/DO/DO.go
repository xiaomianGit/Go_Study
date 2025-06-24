package DO

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id        uint `gorm:"primarykey;auto_increment"`
	CreatedAt time.Time
	Created   string `gorm:varchar(20)"`
	UpdatedAt time.Time
	Updated   string `gorm:varchar(20)"`
}

// SQL题目1
type Students struct {
	Base
	Name  string `gorm:"varchar(20)"`
	Age   uint   `gorm:"int(4)"`
	Grade string `gorm:"varchar(20)"`
}

// SQL题目2
type Accounts struct {
	Base
	Name    string  `gorm:"varchar(20)"`
	Balance float64 `gorm:"float(11,2)"`
}
type Transactions struct {
	Base
	FromAccountId uint    `gorm:"int(11)"`
	ToAccountId   uint    `gorm:"int(11)"`
	Amount        float64 `gorm:"float(11,2)"`
}

// SQLX题目
type Employees struct {
	Base
	Name       string  `gorm:"varchar(20)"`
	Department string  `gorm:"varchar(20)"`
	Salary     float64 `gorm:"float(11,2)"`
}
type Books struct {
	Base
	Title  string  `gorm:"varchar(20)"`
	Author string  `gorm:"varchar(20)"`
	Price  float64 `gorm:"float(11,2)"`
}

// high
type User struct {
	Base
	Username string `gorm:"varchar(20)"`
	Posts    []Post `gorm:"foreignkey:UserId"`
}
type Post struct {
	Base
	Title      string `gorm:"varchar(20)"`
	Content    string `gorm:"varchar(20)"`
	CommentNum uint   `gorm:"int(11)"`
	Status     string `gorm:"varchar(20)"`
	UserId     uint   `gorm:"int(11)"`
	User       User   `gorm:"foreignkey:UserId"`
}
type Comment struct {
	Base
	Content string `gorm:"varchar(20)"`
	PostId  uint   `gorm:"int(11)"`
	UserId  uint   `gorm:"int(11)"`
	Post    Post   `gorm:"foreignkey:PostId"`
	User    User   `gorm:"foreignkey:UserId"`
}

func (p *Post) AfterUpdate(db *gorm.DB) error {
	if skip, exists := db.Get("skip_after_update"); exists && !skip.(bool) {
		fmt.Printf("[afterUpdate] skip_after_update")
		return nil // 跳过钩子
	}
	fmt.Printf("[afterUpdate]")
	var commentNum int64
	db.Model(Comment{}).Where("post_id = ?", p.Id).Count(&commentNum)

	/*result := db.Model(&Post{}).Where("id", p.Id).Update("comment_num", commentNum).Error
	if result.Error != nil {
		panic(result.Error)
	}*/
	fmt.Printf("[afterUpdate] now commentNum is :%d\n", commentNum)

	// 执行后移除标志（可选，增强安全性）
	db.Statement.Context = context.WithValue(db.Statement.Context, "skip_after_update", false)
	return nil
}

func (c *Comment) AfterDelete(db *gorm.DB) error {
	fmt.Printf("[afterDelete] %v \n", c)
	var commentNum int64
	db.Model(&Comment{}).Where("post_id = ?", c.PostId).Count(&commentNum)

	var status string
	if commentNum == 0 {
		status = "无评论"
	} else {
		status = "有评论"
	}
	result := db.Model(&Post{}).Where("id = ?", c.PostId).Updates(map[string]interface{}{
		"status":     status,
		"CommentNum": commentNum,
	})
	if result.Error != nil {
		panic(result.Error)
	}
	fmt.Printf("[afterDelete] after commentNum is :%d\n", commentNum)
	return nil
}
