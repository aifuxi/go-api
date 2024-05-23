package api

import (
	"fmt"
	"github.com/aifuxi/go-api/db/sqlc"
	"github.com/aifuxi/go-api/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries    *sqlc.Queries
	router     *gin.Engine
	tokenMaker token.Maker
}

const (
	secretKey = "abcdefghijklmnop"
)

func NewServer(queries *sqlc.Queries) (*Server, error) {
	maker, err := token.NewJWTMaker(secretKey)
	if err != nil {
		return nil, fmt.Errorf("could not create token maker: %w", err)
	}

	router := gin.Default()
	server := &Server{queries: queries, router: router, tokenMaker: maker}

	router.GET("/users", server.listUsers)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.login)

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
