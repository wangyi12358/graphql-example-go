package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/graph"
	"go-gin-example/internal/middleware"
	"go-gin-example/internal/router"
	"go-gin-example/internal/service/lov"
	"go-gin-example/internal/service/lov_field"
	"go-gin-example/internal/service/user"
	"go-gin-example/pkg/config"
	"go-gin-example/pkg/model"
)

func init() {
	config.Setup()
	model.Setup()
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService:     user.New(model.DB),
		LovService:      lov.New(model.DB),
		LovFieldService: lov_field.New(model.DB),
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(corsConfig))

	r.Use(middleware.AuthMiddleware())

	// init api
	router.InitRouter(r)

	// graphql
	r.Use(graph.GinContextToContextMiddleware())
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
