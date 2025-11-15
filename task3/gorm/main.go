package main

/* 假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。 */

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     		string
	Password 		string
	Posts    		[]Post
	PostCount      int `gorm:"default:0"` // 文章数量统计字段
}

type Post struct { 
	gorm.Model
	Name 			string
	Title   		string
	Content 		string
	UserID  		uint // 外键关联到 User
	User    		User `gorm:"foreignKey:UserID"` // 反向关联
	Comments 		[]Comment
	PostCount      	int `gorm:"default:0"` // 文章数量统计字段
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint // 外键关联到 Post
	Post    Post `gorm:"foreignKey:PostID"` // 反向关联
	UserID  uint // 外键关联到 User（可选：评论者）
	User    User `gorm:"foreignKey:UserID"` // 反向关联（可选）
}

// 查询指定用户的所有文章及对应评论
func QueryUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
    var posts []Post
    err := db.Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
    return posts, err
}

// 查询评论数量最多的文章
func QueryPostWithMostComments(db *gorm.DB) (*Post, error) {
    var post Post
    err := db.Table("posts").
        Select("posts.*, COUNT(comments.id) as comment_count").
        Joins("left join comments on posts.id = comments.post_id").
        Group("posts.id").
        Order("comment_count DESC").
        Limit(1).
        Find(&post).Error
    
    return &post, err
}

// Post 模型的创建钩子函数
func (p *Post) AfterCreate(tx *gorm.DB) error {
    // 更新用户的文章数量统计
    return tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

// Comment 模型的删除钩子函数
func (c *Comment) AfterDelete(tx *gorm.DB) error {
    // 检查文章的评论数量
    var commentCount int64
    tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)
    
    // 如果评论数量为0，则更新文章的评论状态
    if commentCount == 0 {
        return tx.Model(&Post{}).Where("id = ?", c.PostID).UpdateColumn("comment_status", "无评论").Error
    }
    
    return nil
}

func main() { 
	dsn := "root:123456@tcp(192.168.100.121:3306)/task3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	if err != nil {
		fmt.Println("数据库连接失败") // 打印错误信息
	}

	// 迁移表
	db.AutoMigrate(&User{}, &Post{}, &Comment{})   // 题目一：创建数据库表

	// 题目二关联查询
	// 查询ID为1的用户发布的所有文章及评论
	/* if posts, err := QueryUserPostsWithComments(db, 1); err != nil {
        fmt.Println("查询失败:", err)
    } else {
        fmt.Printf("用户文章数量: %d\n", len(posts))
        for _, post := range posts {
            fmt.Printf("文章: %s, 评论数: %d\n", post.Title, len(post.Comments))
        }
	} */

	 // 查询评论数最多的文章
	/* if post, err := QueryPostWithMostComments(db); err != nil {
        fmt.Println("查询失败:", err)
    } else {
        fmt.Printf("评论数最多的文章: %s\n", post.Title)
    } */


}