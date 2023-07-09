package testmain

import "github.com/jackc/pgx/v4"

type Service struct {
	db *pgx.Conn
}

type Product struct {
	Name  string
	Price int
}

func (s *Service) Products() []Product {
	data := []Product{
		{
			Name:  "Компьютер",
			Price: 20_000,
		},
	}
	return data
}
