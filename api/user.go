package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/aifuxi/go-api/db/sqlc"
	"github.com/aifuxi/go-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserRequest struct {
	Username       string `form:"username" binding:"required"`
	HashedPassword string `form:"hashedPassword" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	params := sqlc.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}

	user, err := server.queries.CreateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

type listUsersRequest struct {
	Offset int64 `form:"offset" binding:"min=0"`
	Limit  int64 `form:"limit" binding:"required,min=10,max=50"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(req)
	params := sqlc.ListUsersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	users, err := server.queries.ListUsers(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

type loginRequest struct {
	Username       string `form:"username" binding:"required"`
	HashedPassword string `form:"hashedPassword" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := server.queries.GetUser(ctx, req.Username)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect username or password"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pass := util.CheckPasswordHash(req.HashedPassword, user.HashedPassword)
	if !pass {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect username or password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "success"})
}
