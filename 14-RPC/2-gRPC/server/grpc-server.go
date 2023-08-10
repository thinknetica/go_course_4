package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-core-4/14-RPC/2-gRPC/books_proto"
)

type Bookinist struct {
	Data []pb.Book

	// Композиция типов.
	pb.UnimplementedBookinistServer
}

func (b *Bookinist) Books(_ *pb.Empty, stream pb.Bookinist_BooksServer) error {
	for i := range b.Data {
		stream.Send(&b.Data[i])
	}
	return nil
}

func (b *Bookinist) AddBook(_ context.Context, book *pb.Book) (*pb.Empty, error) {
	b.Data = append(b.Data, *book)
	return new(pb.Empty), nil
}

func main() {
	srv := Bookinist{}
	srv.Data = append(srv.Data,
		pb.Book{Id: 2, Title: "The Go Programming Language"},
		pb.Book{Id: 1, Title: "1984"},
	)

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookinistServer(grpcServer, &srv)
	grpcServer.Serve(lis)
}
