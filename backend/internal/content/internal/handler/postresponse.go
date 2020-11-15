package handler

import "time"

type PostResponse struct {
	Date              time.Time
	DateGMT           time.Time
	GUID              string
	ID                string
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
	Author            int
	Excerpt           string
	FeaturedMedia     int
	CommentStatus     string
	PingStatus        string
	Format            string
	Meta              string
	Sticky            bool
	Template          string
	Categories        []int
	Tags              []int
}
