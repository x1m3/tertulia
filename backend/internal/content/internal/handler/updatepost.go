package handler

import (
	"github.com/gin-gonic/gin"
	ulid "github.com/x1m3/Tertulia/backend/pkg/id"
)

type updatePostRequest struct {
	ID ulid.ID
	newPostRequest
}

func UpdatePost(ctx *gin.Context) (resp interface{}, err error) {
	var req updatePostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{ID: ulid.New()}, nil
}
