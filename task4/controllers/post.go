package controllers

import (
	"task4/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// 定义结构体，用于接收文章标题及内容数据
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}
	// 绑定到上面的结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}

	// 创建一个接受文章的结构体
	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}
	// 将文章添加到数据库
	if err := models.DB.Create(&post).Error; err != nil {
		c.JSON(500, gin.H{"error": "创建文章失败"})
		return
	}
	c.JSON(200, gin.H{"message": "创建文章成功"})
}

func GetPosts(c *gin.Context) {
	// 获取文章ID
	id := c.Param("id")
	var post models.Post
	// 查找指定指定id的文章信息
	if err := models.DB.Preload("User").First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}
	c.JSON(200, gin.H{"post": post})
}

func UpdatePost(c *gin.Context) {
	// 获取用户ID和文章ID
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	// 查找对应id的文章
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}
	// 验证要修改的文章是否属于当前登录用户
	if post.UserID != userID {
		c.JSON(403, gin.H{"error": "您无权修改该文章"})
		return
	}
	// 定义一个接受文章的结构体
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	// 绑定上面的结构体
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 更新数据库文章内容
	models.DB.Model(&post).Updates(input)
	c.JSON(200, gin.H{"message": "更新文章成功"})
}

func DeletePost(c *gin.Context) {
	// 获取用户ID和文章ID
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	// 获取文章信息
	var post models.Post
	if err := models.DB.First(&post, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "文章不存在"})
		return
	}
	// 验证要删除的文章是否属于当前登录用户
	if post.UserID != userID {
		c.JSON(403, gin.H{"error": "您无权删除该文章"})
		return
	}
	// 删除文章
	models.DB.Delete(&post)
	c.JSON(200, gin.H{"message": "删除文章成功"})
}
