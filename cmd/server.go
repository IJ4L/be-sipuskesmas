package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IJ4L/internal/graph"
	resolver "github.com/IJ4L/internal/graph/resolvers"
	"github.com/IJ4L/internal/injector"
	"github.com/IJ4L/pkg/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	cfg, err := utils.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	resolver, err := injector.InitializeResolver(cfg)
	if err != nil {
		log.Fatalf("failed to initialize resolver: %v", err)
	}

	graphQLServe(cfg, resolver)
}

func graphQLServe(cfg utils.Config, rsl *resolver.Resolver) {
	r := gin.Default()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{UserUsecase: rsl.UserUsecase}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	r.POST("/query", gin.WrapH(srv))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))

	log.Fatal(r.Run(":" + cfg.App))
}
