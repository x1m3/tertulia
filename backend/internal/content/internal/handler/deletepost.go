package handler

import (
	"github.com/gin-gonic/gin"
	ulid "github.com/x1m3/Tertulia/backend/pkg/id"
)

type deletePostRequest struct {
	ID    ulid.ID
	Force bool
}

func DeletePost(ctx *gin.Context) (resp interface{}, err error) {
	var req deletePostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{ID: ulid.New()}, nil
}
