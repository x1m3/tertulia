package domain

import (
	"time"

	"github.com/x1m3/tertulia/internal/pkg/id"
)

// Topic represents a forum post
type Topic struct {
	ID        id.ID
	Title     string
	Summary   string
	URL       *string
	Image     *string
	Video     *string
	AuthorID  id.ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Comment represents a comment on a forum post
type Comment struct {
	ID              id.ID
	Content         string
	AuthorID        id.ID
	TopicID         id.ID
	ParentCommentID *id.ID
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// User represents a user in the system
type User struct {
	ID        id.ID
	Username  string
	Email     string
	FirstName string
	LastName  string
	Bio       string
	AvatarURL *string
	LastLogin *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
