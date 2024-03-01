package router

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/graph"
	"go-gin-example/internal/graph/graph_model"
	"go-gin-example/internal/middleware"
	"go-gin-example/internal/service/lov"
	"go-gin-example/internal/service/lov_field"
	"go-gin-example/internal/service/user"
	"go-gin-example/pkg/model"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := graph.Config{Resolvers: &graph.Resolver{
		UserService:     user.New(model.DB),
		LovService:      lov.New(model.DB),
		LovFieldService: lov_field.New(model.DB),
	}}
	c.Directives.Constraint = func(ctx context.Context, obj interface{}, next graphql.Resolver, format *string, name *string) (res interface{}, err error) {
		return next(ctx)
	}
	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, requires *graph_model.Role) (res interface{}, err error) {
		sysUser := middleware.ForContext(ctx)
		if sysUser != nil {
			fmt.Printf("id: %d", sysUser.ID)
		}
		return next(ctx)
	}
	h := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func InitGraph(r *gin.Engine) {
	//query := r.Group("/query")
	//query.Use(middleware.AuthMiddleware())
	r.Use(middleware.AuthMiddleware()).POST("/query", graphqlHandler())
}
