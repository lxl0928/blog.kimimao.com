package main

import (
	"log"

	"blog.kimimao.com/backend/config"
	"blog.kimimao.com/backend/handler"
	"blog.kimimao.com/backend/pkg/database"
	"blog.kimimao.com/backend/pkg/redis"
	"blog.kimimao.com/backend/repository"
	"blog.kimimao.com/backend/router"
	"blog.kimimao.com/backend/service"
)

func main() {
	cfg := config.Load()

	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	defer db.Close()

	if err := redis.InitRedis(&cfg.Redis); err != nil {
		log.Printf("Redis连接失败(将跳过缓存): %v", err)
	}

	// 仓库层
	userRepo := repository.NewUserRepository(db)
	blogRepo := repository.NewBlogRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	likeRepo := repository.NewLikeFavoriteRepository(db)

	// 服务层
	userService := service.NewUserService(userRepo, cfg)
	blogService := service.NewBlogService(blogRepo, likeRepo)
	commentService := service.NewCommentService(commentRepo, blogRepo)

	// 处理器
	uh := handler.NewUserHandler(userService)
	bh := handler.NewBlogHandler(blogService)
	ch := handler.NewCommentHandler(commentService)

	r := router.Setup(cfg, uh, bh, ch)
	log.Printf("服务启动: http://:%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
