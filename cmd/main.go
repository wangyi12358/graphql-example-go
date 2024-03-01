package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-gin-example/internal/graph"
	"go-gin-example/internal/router"
	"go-gin-example/pkg/config"
	"go-gin-example/pkg/model"
	"go-gin-example/pkg/validate"
)

func init() {
	config.Setup()
	model.Setup()
	validate.Setup()
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

	r.Use(graph.GinContextToContextMiddleware())

	// init api
	router.InitRouter(r)

	// init graphql
	router.InitGraph(r)

	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%d", config.Config.Server.Port))
}
