package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go-gin-example/internal/model/sys_user_model"
	"go-gin-example/pkg/token"
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
		tokenStr := c.GetHeader("authorization")
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		if tokenStr != "" {
			id, _ := token.ParseToken(tokenStr)
			fmt.Printf("id: %d", id)
			if id != nil {
				user, _ := sys_user_model.FindById(*id)
				fmt.Printf("id: %d", user.ID)
				c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), userCtxKey, user))
			}
		}
		c.Next()
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *sys_user_model.SysUser {
	raw, _ := ctx.Value(userCtxKey).(*sys_user_model.SysUser)
	return raw
}
