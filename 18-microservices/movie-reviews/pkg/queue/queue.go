package queue

import "go-core-4/18-microservices/movie-reviews/pkg/reviews"

// Interface - контракт драйвера для работы с очередью сообщений.
type Interface interface {
	Publish(queue string, rivew reviews.Review) error
	Fetch(queue string) (reviews.Review, error)
}
