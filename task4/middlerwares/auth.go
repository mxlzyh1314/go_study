package middlerwares

import (
	"strings"
	"task4/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	// JWTAuthMiddleware 是一个 Gin 中间件函数，用于验证请求中的 JWT token
	// 返回值: gin.HandlerFunc - Gin 处理器函数
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段作为 token
		authHeader := c.GetHeader("Authorization")
		// 检查 token 是否存在，如果不存在则返回 401 未授权错误
		if authHeader == "" {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "未授权",
			})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(401, gin.H{
				"code":    401,
				"message": "需要token",
			})
			c.Abort()
			return
		}

		// 解析JWT token并验证用户身份
		// 如果token无效则返回401无效的token并终止请求处理
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			// token解析失败，返回未授权错误
			c.JSON(401, gin.H{
				"code":    401,
				"message": "无效的token",
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
