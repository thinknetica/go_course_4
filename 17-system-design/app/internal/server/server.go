package server

import (
	"go-core-4/17-system-design/app/internal/api"
	"go-core-4/17-system-design/app/internal/db"
)

type Server struct {
	api *api.API
	db  *db.DB
}

const connStr = "postgres://..."

func New() (*Server, error) {
	s := Server{}
	db, err := db.New(connStr)
	if err != nil {
		return nil, err
	}
	s.db = db
	s.api = api.New(s.db)

	return &s, nil
}

func (s *Server) Run() {
	// ...
}
