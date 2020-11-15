package handler

import (
	ulid "github.com/x1m3/Tertulia/backend/pkg/id"
	"time"
)

type PostResponse struct {
	Date              time.Time
	DateGMT           time.Time
	GUID              ulid.ID
	ID                ulid.ID
	Link              string
	Modified          time.Time
	ModifiedGMT       time.Time
	Slug              string
	Status            string
	Type              string
	Password          string
	PermalinkTemplate string
	GeneratedSlug     string
	Title             string
	Content           string
	Author            ulid.ID
	Excerpt           string
	FeaturedMedia     ulid.ID
	CommentStatus     string
	PingStatus        string
	Format            string
	Meta              string
	Sticky            bool
	Template          string
	Categories        []ulid.ID
	Tags              []ulid.ID
}
