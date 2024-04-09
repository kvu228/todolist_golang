package middlewares

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"to_do_list/common"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (requester common.Requester, err error)
}

func RequireAuth(ac AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		requester, err := ac.IntrospectToken(c.Request.Context(), token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Set(common.KeyRequester, requester)
		c.Next()
	}
}

func extractTokenFromHeaderString(header string) (string, error) {
	// "Authorization": "Bearer {token}"
	parts := strings.Split(header, " ")
	// parts[0] = "Bearer"
	// parts[1] = "{token}"
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("missing access token")
	}
	return parts[1], nil
}
