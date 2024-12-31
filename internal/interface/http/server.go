package http

import (
	"context"

	"github.com/x1m3/tertulia/internal/pkg/config"
)

// Server is a struct that implements the API interface
type Server struct {
	cfg *config.Config
}

// DeleteComment deletes a comment
// DELETE /comments/{comment_id}
func (s Server) DeleteComment(ctx context.Context, request DeleteCommentRequestObject) (DeleteCommentResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// GetComment returns a comment
// GET /comments/{comment_id}
func (s Server) GetComment(ctx context.Context, request GetCommentRequestObject) (GetCommentResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateComment updates a comment
// PUT /comments/{comment_id}
func (s Server) UpdateComment(ctx context.Context, request UpdateCommentRequestObject) (UpdateCommentResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// ListTopics returns a list of topics
// GET /topics
func (s Server) ListTopics(ctx context.Context, request ListTopicsRequestObject) (ListTopicsResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// CreateTopic creates a topic
// POST /topics
func (s Server) CreateTopic(ctx context.Context, request CreateTopicRequestObject) (CreateTopicResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// DeleteTopic deletes a topic
// DELETE /topics/{topic_id}
func (s Server) DeleteTopic(ctx context.Context, request DeleteTopicRequestObject) (DeleteTopicResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// GetTopic returns a topic
// GET /topics/{topic_id}
func (s Server) GetTopic(ctx context.Context, request GetTopicRequestObject) (GetTopicResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// UpdateTopic updates a topic
// PUT /topics/{topic_id}
func (s Server) UpdateTopic(ctx context.Context, request UpdateTopicRequestObject) (UpdateTopicResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// ListComments returns a list of comments
// GET /comments
func (s Server) ListComments(ctx context.Context, request ListCommentsRequestObject) (ListCommentsResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// CreateComment creates a comment
// POST /comments
func (s Server) CreateComment(ctx context.Context, request CreateCommentRequestObject) (CreateCommentResponseObject, error) {
	// TODO implement me
	panic("implement me")
}

// HealthCheck returns info about the server health status and important dependencies
// GET /health
func (s Server) HealthCheck(ctx context.Context, request HealthCheckRequestObject) (HealthCheckResponseObject, error) {
	return HealthCheck200JSONResponse{Http: true}, nil
}

// NewServer is a Server constructor
func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
