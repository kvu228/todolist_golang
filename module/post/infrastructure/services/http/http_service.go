package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/common"
	"to_do_list/middlewares"
	"to_do_list/module/post/usecase"
	"to_do_list/module/post/usecase/command"
	"to_do_list/module/post/usecase/query"
)

type httpPostService struct {
	postQueryUseCase command.PostCmdUseCase
	postCmdUseCase   query.PostQueryUseCase
	authClient       middlewares.AuthClient
}

func NewHttpPostService(postQueryUseCase command.PostCmdUseCase, postCmdUseCase query.PostQueryUseCase) *httpPostService {
	return &httpPostService{postQueryUseCase: postQueryUseCase, postCmdUseCase: postCmdUseCase}
}

func (s *httpPostService) handleListPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param usecase.ListPostsParams
		if err := c.Bind(param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := s.postCmdUseCase.ListPosts(c.Request.Context(), &param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func (s *httpPostService) handleCreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto usecase.NewPostDTO
		if err := c.Bind(&dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		requester := c.MustGet(common.KeyRequester).(common.Requester)
		dto.OwnerId = requester.Id()

		if err := s.postQueryUseCase.CreatePost(c.Request.Context(), dto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})

	}
}

func (s *httpPostService) Routes(g *gin.RouterGroup) {
	post := g.Group("/post")
	{
		post.GET("", s.handleListPosts())
		post.POST("/create", middlewares.RequireAuth(s.authClient), s.handleCreatePost())
	}
}

func (s *httpPostService) SetAuthClient(authClient middlewares.AuthClient) *httpPostService {
	s.authClient = authClient
	return s
}
