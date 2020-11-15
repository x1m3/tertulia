package handler

import (
	"github.com/gin-gonic/gin"
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
	Author        int
	Excerpt       string
	FeaturedMedia int
	CommentStatus string
	PingStatus    string
	Format        string
	Meta          string
	Sticky        bool
	Template      string
	Categories    []int
	Tags          []int
}

func NewPost(ctx *gin.Context) (resp interface{}, err error) {
	var req newPostRequest
	if err := ctx.ShouldBind(&req); err != nil {
		return nil, err
	}
	return &PostResponse{}, nil
}
