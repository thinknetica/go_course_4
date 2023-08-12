package reviews

import "sync"

// Review - рецензия на фильм.
type Review struct {
	MovieID int
	UserID  int
	Review  string
}

type DB struct {
	mu   sync.Mutex
	data []Review
}

func (db *DB) GetAll() []Review {
	db.mu.Lock()
	data := make([]Review, len(db.data))
	copy(data, db.data)
	db.mu.Unlock()
	return data
}

func (db *DB) Add(r Review) {
	db.mu.Lock()
	db.data = append(db.data, r)
	db.mu.Unlock()
}
