package handler

import "github.com/gin-gonic/gin"

type deletePostRequest struct {
	ID    int
	Force bool
}

func DeletePost(ctx *gin.Context) (resp interface{}, err error) {
	var req deletePostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{}, nil
}
