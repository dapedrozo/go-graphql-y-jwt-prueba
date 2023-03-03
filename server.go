package main

import (
	"GoGraphQlJwt/http"
	"GoGraphQlJwt/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	server := gin.Default()
	server.Use(middleware.BasicAuth())
	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", middleware.AuthorizeJWT(), http.GraphQLHandler())

	server.Run(defaultPort)

}
