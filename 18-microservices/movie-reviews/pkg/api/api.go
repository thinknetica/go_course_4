package api

import "go-core-4/18-microservices/movie-reviews/pkg/reviews"

type API struct {
	db *reviews.DB
}

func New(db *reviews.DB) *API {
	api := API{
		db: db,
	}
	return &api
}
