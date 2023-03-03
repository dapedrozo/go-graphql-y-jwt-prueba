package http

import (
	"GoGraphQlJwt/graph"
	"GoGraphQlJwt/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func GraphQLHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(ctx *gin.Context) {
		srv.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
