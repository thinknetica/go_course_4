package main

import (
	"go-core-4/13-api/2-api/pkg/api"
	"go-core-4/18-microservices/movie-reviews/pkg/queue"
	"go-core-4/18-microservices/movie-reviews/pkg/reviews"
)

type Server struct {
	api   *api.API
	queue *queue.Interface
	db    *reviews.DB
}

func New() *Server {
	s := Server{
		api:   api.New(),
		queue: nil,
		db:    nil,
	}
	return &s
}

func main() {
	srv := New()
	srv.run()
}

func (s *Server) run() {

}
