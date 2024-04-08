package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/module/post/domain"
	"to_do_list/module/post/infrastructure/repositories/rgpc_http"
	"to_do_list/module/post/usecases/query"
)

type httpPostService struct {
	postRepository query.PostRepository
}

func NewHttpPostService(postRepository query.PostRepository) *httpPostService {
	return &httpPostService{postRepository: postRepository}
}

func (s *httpPostService) handleListPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var param domain.ListPostsParams
		if err := c.Bind(param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		/*
			The Code below is used for establish rpc connect to http server of catRepo

		*/
		//configComp := s.sctx.MustGet(common.KeyConfig).(interface{ GetURL() string })
		//urlRPC := fmt.Sprintf("%s/query-categories-ids", configComp.GetURL())

		userRepo := rgpc_http.NewRpcGetUsersByIds("http://127.0.0.1:8080/api/v1/user/rpc/query-users-ids")
		result, err := query.NewListPostsQueryUseCase(s.postRepository, userRepo).ListPosts(c.Request.Context(), &param)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}

func (s *httpPostService) Routes(g *gin.RouterGroup) {
	products := g.Group("/products")
	{
		products.GET("", s.handleListPosts())
	}
}
