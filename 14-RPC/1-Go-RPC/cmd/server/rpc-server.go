package main

import (
	"log"
	"net"
	"net/rpc"

	"go-core-4/14-rpc/1-go-rpc/pkg/books"
)

// Server - тип данных RCP-сервера.
type Server int

// Books возвращает список книг.
func (s *Server) Books(req books.Request, resp *[]books.Book) error {
	// resp = &allBooks - не работает.
	*resp = allBooks
	return nil
}

// Book возвращает книгу по ID
func (s *Server) Book(req books.Request, resp *books.Book) error {
	for _, b := range allBooks {
		if b.ID == req.ID {
			*resp = b
			return nil
		}
	}
	return nil
}

func main() {
	srv := new(Server)
	err := rpc.Register(srv)
	if err != nil {
		log.Fatal(err)
	}

	// регистрация сетевой службы RPC-сервера
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// rpc.Accept(listener) // блокирующая команда.
	// цикл обработки клиентских подключений
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go rpc.ServeConn(conn)
	}
}

var allBooks = []books.Book{
	{
		ID:     1,
		Title:  "The Lord Of The Rings",
		Author: "J.R.R. Tolkien",
	},
	{
		ID:     2,
		Title:  "The Chronicles of Amber",
		Author: "Roger Zelazny",
	},
}
