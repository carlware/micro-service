package graphql

import (
	"carlware/accounts/cli/config"
	"carlware/accounts/cli/dispatchers"
	"carlware/accounts/cli/dispatchers/graphql/graph"
	"carlware/accounts/cli/dispatchers/graphql/graph/generated"
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func graphqlHandler(crtl *dispatchers.Controller) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{crtl}}))

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

func NewGraphQL(ctx context.Context, cfg *config.Configuration, crtl *dispatchers.Controller) error {
	// Setting up Gin
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))
	r.POST("/query", graphqlHandler(crtl))
	r.GET("/", playgroundHandler())
	return r.Run(cfg.GraphQL.Host + ":" + cfg.GraphQL.Port)
}
