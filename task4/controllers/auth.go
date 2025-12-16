package controllers

import (
	"task4/models"
	"task4/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 注册用户
func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// 创建一个新的用户对象
	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
		Email:    input.Email,
	}
	// 将用户数据保存到数据库中
	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(200, gin.H{"message": "注册成功"})
}

// 登录
func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	// 验证输入
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// 验证登录用户名
	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "用户名错误"})
		return
	}
	// 验证登录用户名的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "密码错误"})
		return
	}

	// 生成JWT
	token, err := utils.GenerateJwt(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "登录失败"})
		return
	}
	c.JSON(200, gin.H{"token": token})

}
