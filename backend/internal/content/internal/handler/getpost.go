package handler

import (
	"github.com/gin-gonic/gin"
	ulid "github.com/x1m3/Tertulia/backend/pkg/id"
)

type getPostRequest struct {
	ID       ulid.ID
	Context  string
	Password string
}

func GetPost(ctx *gin.Context) (resp interface{}, err error) {
	var req getPostRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{ID: ulid.New()}, nil
}
