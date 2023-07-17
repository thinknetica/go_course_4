package dip

// База данных - деталь реализации.
type DB []Book

type Book struct {
	ID     int
	Title string
}

// Интерфейс - часть бизнес-логики.
type Storage interface {
	Books() []Book
}

func (db DB) Books() []Book {
	return []Book{}
}

// Нарушение принципа DIP.
type WrongServer struct {
	db DB
}

// Выполнение принципа DIP.
type GoodServer struct {
	db Storage
}

// Бизнес-логика системы.
func (s *WrongServer) BusinessLogic() {
	// Принцип DIP нарушен, обращение к БД.
	books := s.db.Books()
	_ = books
}

// Бизнес-логика системы.
func (s *GoodServer) BusinessLogic() {
	// Принцип DIP выполняется, обращение к интерфейсу.
	books := s.db.Books()
	_ = books
}
