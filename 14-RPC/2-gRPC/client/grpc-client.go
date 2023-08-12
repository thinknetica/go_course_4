package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "go-core-4/14-rpc/2-grpc/books_proto"
)

func main() {
	conn, err := grpc.Dial("localhost:12345",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewBookinistClient(conn)

	ctx := context.Background()

	err = printAllBooksOnserver(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	client.AddBook(context.Background(), &pb.Book{Id: 3, Title: "The Lord Of The Rings"})

	err = printAllBooksOnserver(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

func printAllBooksOnserver(ctx context.Context, client pb.BookinistClient) error {
	fmt.Println("\nЗапрашиваю книги на gRPC-сервере.")
	stream, err := client.Books(context.Background(), &pb.Empty{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			book, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			fmt.Printf("Получена книга: %v\n", book)
		}
	}
}
