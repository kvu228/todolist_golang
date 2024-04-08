package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"to_do_list/module/users/usecase"
	"to_do_list/module/users/usecase/command"
	"to_do_list/module/users/usecase/query"
)

type httpUserService struct {
	userQueryUseCase query.UserQueryUseCase
	userCmdUseCase   command.UserCmdUseCase
}

func NewHttpUserService(userQueryUseCase query.UserQueryUseCase, userCmdUseCase command.UserCmdUseCase) *httpUserService {
	return &httpUserService{
		userQueryUseCase: userQueryUseCase,
		userCmdUseCase:   userCmdUseCase,
	}
}

func (s *httpUserService) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		user, err := s.userQueryUseCase.GetUser(c.Request.Context(), id)
		if err != nil {
			c.JSON(200, gin.H{
				"Error": err.Error(),
			})
			return
		}

		c.JSON(
			http.StatusOK, gin.H{
				"data": user,
			})
	}
}

func (s *httpUserService) listUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params struct {
			IdsStr []string `json:"ids"`
		}

		if err := c.BindJSON(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		Ids := make([]uuid.UUID, len(params.IdsStr))
		for i, id := range params.IdsStr {
			Ids[i], _ = uuid.Parse(id)
		}

		result, err := s.userQueryUseCase.ListUsersByIds(c.Request.Context(), Ids)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

func (s *httpUserService) handleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordRegistrationDTO
		if err := c.BindJSON(&dto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := s.userCmdUseCase.Register(c.Request.Context(), dto); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(
			http.StatusOK, gin.H{
				"data": true,
			})
	}
}

func (s *httpUserService) handleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.EmailPasswordLoginDTO
		if err := c.BindJSON(&dto); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.userCmdUseCase.LoginEmailPassword(c.Request.Context(), dto)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})

	}
}

func (s *httpUserService) Routes(g *gin.RouterGroup) {
	g.POST("/register", s.handleRegister())
	g.POST("/authenticate", s.handleLogin())

	user := g.Group("/user")
	user.GET("/:id", s.getUser())
	user.POST("/", s.listUsers())

	rpc := user.Group("/rpc")
	rpc.POST("/query-users-ids", s.listUsers())
}
