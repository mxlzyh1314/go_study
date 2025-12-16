package controllers

import (
	"strconv"
	"task4/models"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	// 获取用户ID和文章ID
	userID := c.MustGet("userID").(uint)
	postID := c.Param("id")

	// 创建结构体接受数据
	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	postIDUint, _ := strconv.ParseUint(postID, 10, 32)
	comment := models.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  uint(postIDUint),
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"comment": comment})
}

func GetComments(c *gin.Context) {
	// 获取文章ID
	postID := c.Param("id")

	// 查询所有评论
	var comments []models.Comment
	if err := models.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"comments": comments})
}
