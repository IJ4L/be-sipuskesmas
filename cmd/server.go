package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/IJ4L/internal/graph"
	resolver "github.com/IJ4L/internal/graph/resolvers"
	"github.com/IJ4L/pkg/db/postgres"
	"github.com/IJ4L/pkg/utils"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	cfg, err := utils.LoadConfig(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	_, err = postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	GraphQLServe(cfg)
}

func GraphQLServe(cfg utils.Config) {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.App)
	log.Fatal(http.ListenAndServe(":"+cfg.App, nil))
}
