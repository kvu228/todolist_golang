package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"to_do_list/common"
	components "to_do_list/component"
	PostMySQL "to_do_list/module/post/infrastructure/repositories/mysql"
	PostHTTP "to_do_list/module/post/infrastructure/services/http"
	UserMySQL "to_do_list/module/users/infrastructure/repositories/mysql"
	UserHTTP "to_do_list/module/users/infrastructure/services/http"
	"to_do_list/module/users/usecase/command"
	"to_do_list/module/users/usecase/query"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	db.Debug()

	server := gin.Default()
	v1 := server.Group("/api/v1")
	secretKey := os.Getenv("JWT_SECRET")
	tokenProvider := components.NewJWTProvider(secretKey, 60*60*24*7, 60*60*24*14)
	userRepo := UserMySQL.NewUserMySQLRepo(db)
	userQueryUC := query.NewUserQueryUseCase(userRepo)
	sessionRepo := UserMySQL.NewSessionMySQLRepo(db)
	userCmdUC := command.NewUserCmdUseCase(userRepo, sessionRepo, tokenProvider, &common.Hasher{})
	UserHTTP.NewHttpUserService(userQueryUC, userCmdUC).Routes(v1)

	postRepo := PostMySQL.NewPostsMySQLRepo(db)
	PostHTTP.NewHttpPostService(postRepo).Routes(v1)
	server.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
