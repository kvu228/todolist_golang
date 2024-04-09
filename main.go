package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"to_do_list/common"
	components "to_do_list/component"
	"to_do_list/middlewares"
	PostMySQLRepo "to_do_list/module/post/infrastructure/repositories/mysql"
	PostHTTPService "to_do_list/module/post/infrastructure/services/http"
	UserMySQLRepo "to_do_list/module/users/infrastructure/repositories/mysql"
	UserHTTPService "to_do_list/module/users/infrastructure/services/http"
	UserCmdUC "to_do_list/module/users/usecase/command"
	UserQueryUC "to_do_list/module/users/usecase/query"
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

	//Setup Dependencies
	secretKey := os.Getenv("JWT_SECRET")
	expTokenIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	expRefreshTokenIn, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_TIME"))

	tokenProvider := components.NewJWTProvider(secretKey, expTokenIn, expRefreshTokenIn)
	userRepo := UserMySQLRepo.NewUserMySQLRepo(db)
	userQueryUC := UserQueryUC.NewUserQueryUseCase(userRepo)
	sessionRepo := UserMySQLRepo.NewSessionMySQLRepo(db)
	authClient := UserQueryUC.NewIntrospectUC(userRepo, sessionRepo, tokenProvider)
	userCmdUC := UserCmdUC.NewUserCmdUseCase(userRepo, sessionRepo, tokenProvider, &common.Hasher{})
	UserHTTPService.NewHttpUserService(userQueryUC, userCmdUC).Routes(v1)

	postRepo := PostMySQLRepo.NewPostsMySQLRepo(db)
	PostHTTPService.NewHttpPostService(postRepo).Routes(v1)

	server.GET("ping", middlewares.RequireAuth(authClient), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
