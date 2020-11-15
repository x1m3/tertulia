package handler

import (
	"github.com/gin-gonic/gin"
	"time"
)

// TODO: Review all this wordpress shit. Try to investigate wp types like author, categories, etc.. (numbers, strings)
type listPostsRequest struct {
	Context           string    `json:"context" form:"context"`
	Page              int       `json:"page" form:"page"`
	PerPage           int       `json:"per_page" form:"per_page"`
	Search            string    `json:"search" form:"search"`
	After             time.Time `json:"after" form:"after"`
	Author            string    `json:"author" form:"author"`
	AuthorExclude     string    `json:"author_exclude" form:"author_exclude"`
	Before            time.Time `json:"before" form:"begore"`
	Exclude           []int     `json:"exclude" form:"exclude"`
	Include           []int     `json:"include" form:"include"`
	Offset            int       `json:"offset" form:"offset"`
	Order             string    `json:"order" form:"order"`
	OrderBy           string    `json:"order_by" form:"order_by"`
	Slug              string    `json:"slug" form:"slug"`
	Status            string    `json:"status" form:"status"`
	TaxRelation       string    `json:"tax_relation" form:"tax_relation"`
	Categories        []int     `json:"categories" form:"categories"`
	CategoriesExclude []int     `json:"categories_exclude" form:"categories_exclude"`
	Tags              []int     `json:"tags" form:"tags"`
	TagsExclude       []int     `json:"tags_exclude" form:"tags_exclude"`
	Sticky            bool      `json:"sticky" form:"sticky"`
}

func ListPosts(ctx *gin.Context) (resp interface{}, err error) {
	var req listPostsRequest
	if err := ctx.Bind(&req); err != nil {
		return nil, err
	}
	return []PostResponse{{}}, nil
}
