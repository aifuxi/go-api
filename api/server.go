package api

import (
	"github.com/aifuxi/go-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *sqlc.Queries
	router  *gin.Engine
}

func NewServer(queries *sqlc.Queries) *Server {
	router := gin.Default()
	server := &Server{queries: queries, router: router}

	router.GET("/users", server.listUsers)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.login)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
