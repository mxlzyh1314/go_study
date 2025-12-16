package routers

import (
	"task4/controllers"
	"task4/middlerwares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	public := r.Group("/api")
	{
		// 用户认证路由
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)

		// 文章路由(公开)
		public.GET("/posts", controllers.GetPosts)
		public.GET("/posts/:id", controllers.GetPosts)

	}
	// 受保护路由
	protected := r.Group("/api")
	protected.Use(middlerwares.AuthMiddleware())
	{
		// 文章路由
		protected.POST("/posts", controllers.CreatePost)
		protected.PUT("/posts/:id", controllers.UpdatePost)
		protected.DELETE("/posts/:id", controllers.DeletePost)
		// 评论路由
		protected.POST("/comments", controllers.CreateComment)
		protected.GET("/comments/:id", controllers.GetComments)
	}
	return r
}
