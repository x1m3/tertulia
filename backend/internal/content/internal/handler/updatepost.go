package handler

import "github.com/gin-gonic/gin"

type updatePostRequest struct {
	ID int
	newPostRequest
}

func UpdatePost(ctx *gin.Context) (resp interface{}, err error) {
	var req updatePostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{}, nil
}
