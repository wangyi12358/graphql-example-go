package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go-gin-example/internal/model/sys_user_model"
	"net/http"
)

var userCtxKey = "user"

type GraphqlResponse struct {
	Errors []*gqlerror.Error `json:"errors,omitempty"`
}

func (r *GraphqlResponse) NewError(message string) {
	r.Errors = append(r.Errors, gqlerror.Errorf(message))
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("authorization")
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		res := GraphqlResponse{}
		if token == "" {
			res.NewError("Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *sys_user_model.SysUser {
	raw, _ := ctx.Value(userCtxKey).(*sys_user_model.SysUser)
	return raw
}
