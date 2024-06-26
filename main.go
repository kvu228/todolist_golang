package main

import (
	"fmt"
	ImageUC "to_do_list/module/image/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"to_do_list/common"
	components "to_do_list/component"
	"to_do_list/middlewares"
	ImageMySQLRepo "to_do_list/module/image/infrastructure/repository/mysql"
	ImageHTTPService "to_do_list/module/image/infrastructure/service"
	PostMySQLRepo "to_do_list/module/post/infrastructure/repositories/mysql"
	"to_do_list/module/post/infrastructure/repositories/rgpc_http"
	PostHTTPService "to_do_list/module/post/infrastructure/services/http"
	PostCmdUC "to_do_list/module/post/usecase/command"
	PostQueryUC "to_do_list/module/post/usecase/query"
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

	// user dependencies
	tokenProvider := components.NewJWTProvider(secretKey, expTokenIn, expRefreshTokenIn)
	userRepo := UserMySQLRepo.NewUserMySQLRepo(db)
	userQueryUC := UserQueryUC.NewUserQueryUseCase(userRepo)
	sessionRepo := UserMySQLRepo.NewSessionMySQLRepo(db)
	authClient := UserQueryUC.NewIntrospectUC(userRepo, sessionRepo, tokenProvider)

	// post dependencies
	postRepo := PostMySQLRepo.NewPostsMySQLRepo(db)
	urlUserRPC := fmt.Sprintf("%s/query-users-ids", os.Getenv("URL_RPC_USER"))
	userRPCRepo := rgpc_http.NewRpcGetUsersByIds(urlUserRPC)
	postQueryUC := PostQueryUC.NewPostQueryUseCase(postRepo, userRPCRepo)
	postCmdUC := PostCmdUC.NewPostCmdUseCase(postRepo)
	PostHTTPService.NewHttpPostService(postCmdUC, postQueryUC).SetAuthClient(authClient).Routes(v1)

	// image dependencies
	bucketName := os.Getenv("AWS_S3_BUCKET")
	region := os.Getenv("AWS_S3_REGION")
	s3Api := os.Getenv("AWS_S3_API_KEY")
	s3Secret := os.Getenv("AWS_S3_SECRET")
	cdnDomain := os.Getenv("CDN_DOMAIN")
	s3Uploader := components.NewAWSS3Provider(bucketName, region, s3Api, s3Secret, cdnDomain)
	imageRepo := ImageMySQLRepo.NewImageMySQLRepo(db)
	imageUC := ImageUC.NewImageUseCase(s3Uploader, imageRepo)
	ImageHTTPService.NewHttpImageService(s3Uploader, imageUC).Routes(v1)

	userCmdUC := UserCmdUC.NewUserCmdUseCase(userRepo, sessionRepo, imageRepo, tokenProvider, &common.Hasher{})
	UserHTTPService.NewHttpUserService(userQueryUC, userCmdUC).SetAuthClient(authClient).Routes(v1)

	server.GET("ping", middlewares.RequireAuth(authClient), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.Run(":8080")
}
