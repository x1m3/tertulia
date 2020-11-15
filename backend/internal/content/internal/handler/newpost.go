package handler

import (
	"github.com/gin-gonic/gin"
	ulid "github.com/x1m3/Tertulia/backend/pkg/id"
	"time"
)

type newPostRequest struct {
	Date          time.Time
	DateGMT       time.Time
	Slug          string
	Status        string
	Password      string
	Title         string
	Content       string
	Author        ulid.ID
	Excerpt       string
	FeaturedMedia ulid.ID
	CommentStatus string
	PingStatus    string
	Format        string
	Meta          string
	Sticky        bool
	Template      string
	Categories    []ulid.ID
	Tags          []ulid.ID
}

func NewPost(ctx *gin.Context) (resp interface{}, err error) {
	var req newPostRequest
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{ID: ulid.New()}, nil
}
