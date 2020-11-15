package handler

import "github.com/gin-gonic/gin"

type getPostRequest struct {
	ID       int
	Context  string
	Password string
}

func GetPost(ctx *gin.Context) (resp interface{}, err error) {
	var req getPostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{}, nil
}
