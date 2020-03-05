package http

import (
	"net/http"

	"go-tutorial-2020/pkg/grace"

	"github.com/rs/cors"
)

// UserHandler ...
type UserHandler interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server *http.Server
	User   UserHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
