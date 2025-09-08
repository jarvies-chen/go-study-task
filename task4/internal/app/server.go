package app

import (
	"blog/internal/config"
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/model"
	"blog/internal/repository"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Run(cfg config.Config) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.Source), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)

	// 初始化仓库
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	// 初始化处理函数
	authHandler := handler.NewAuthHandler(userRepo, cfg.JWT.SecretKey, 24*time.Hour)
	postHandler := handler.NewPostHandler(postRepo)
	commentHandler := handler.NewCommentHandler(commentRepo)

	// 初始化路由
	r := gin.Default()

	// 公开路由
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// 受保护路由
	authMiddleware := middleware.AuthMiddleware(cfg.JWT.SecretKey)
	authorized := r.Group("/")
	authorized.Use(authMiddleware)

	{
		authorized.GET("/posts", postHandler.GetPosts)
		authorized.GET("/posts/:id", postHandler.GetPost)
		authorized.POST("/posts", postHandler.CreatePost)
		authorized.PUT("/posts/:id", postHandler.UpdatePost)
		authorized.DELETE("/posts/:id", postHandler.DeletePost)

		authorized.GET("/posts/:id/comments", commentHandler.GetComments)
		authorized.POST("/comments", commentHandler.CreateComment)
	}

	// 启动服务器
	log.Printf("服务器启动在端口 %d", cfg.Server.Port)
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
