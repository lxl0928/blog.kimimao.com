package router

import (
	"blog.kimimao.com/backend/config"
	"blog.kimimao.com/backend/handler"
	"blog.kimimao.com/backend/middleware"
	"github.com/gin-gonic/gin"
)

// Setup 配置路由
func Setup(cfg *config.Config, uh *handler.UserHandler, bh *handler.BlogHandler, ch *handler.CommentHandler) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	// 公开接口
	api := r.Group("/api")
	{
		api.POST("/auth/register", uh.Register)
		api.POST("/auth/login", uh.Login)

		api.GET("/blogs", bh.ListAll)
		api.GET("/blogs/:id", bh.GetDetail)
		api.GET("/blogs/:id/comments", ch.ListByBlogID)
	}

	// 需登录接口
	auth := api.Group("")
	auth.Use(middleware.Auth(cfg))
	{
		auth.GET("/user/profile", uh.GetProfile)
		auth.PUT("/user/profile", uh.UpdateProfile)
		auth.PUT("/user/password", uh.ChangePassword)

		auth.POST("/blogs", bh.Create)
		auth.GET("/blogs/my", bh.ListMy)
		auth.PUT("/blogs/:id", bh.Update)
		auth.DELETE("/blogs/:id", bh.Delete)

		auth.POST("/blogs/:id/comments", ch.Create)
		auth.POST("/blogs/:id/like", bh.Like)
		auth.DELETE("/blogs/:id/like", bh.Unlike)
		auth.POST("/blogs/:id/favorite", bh.Favorite)
		auth.DELETE("/blogs/:id/favorite", bh.Unfavorite)
	}

	// Admin接口
	admin := api.Group("/admin")
	admin.Use(middleware.Auth(cfg), middleware.Admin())
	{
		admin.GET("/users", uh.List)
		admin.GET("/users/stats", uh.GetUserStats)
		admin.GET("/blogs", bh.ListAll)
		admin.GET("/blogs/stats", bh.GetBlogStats)
	}

	return r
}
