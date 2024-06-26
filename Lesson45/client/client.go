package main

import (
    "context"
    "log"

    "google.golang.org/grpc"
	pb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson45/protos"

)

const (
    address = "localhost:50051"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewLibraryServiceClient(conn)

    addBookResponse, err := c.AddBook(context.Background(), &pb.AddBookRequest{
        Title:         "The Catcher in the Rye",
        Author:        "J.D. Salinger",
        YearPublished: 1951,
    })
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
    }
    log.Printf("Added book with ID: %s", addBookResponse.GetBookId())

    searchBookResponse, err := c.SearchBook(context.Background(), &pb.SearchBookRequest{Query: "Catcher"})
    if err != nil {
        log.Fatalf("Could not search book: %v", err)
    }
    log.Println("Search results:")
    for _, book := range searchBookResponse.GetBooks() {
        log.Printf("ID: %s, Title: %s, Author: %s, Year: %d", book.GetBookId(), book.GetTitle(), book.GetAuthor(), book.GetYearPublished())
    }

    borrowBookResponse, err := c.BorrowBook(context.Background(), &pb.BorrowBookRequest{
        BookId: addBookResponse.GetBookId(),
        UserId: "user-123",
    })
    if err != nil {
        log.Fatalf("Could not borrow book: %v", err)
    }
    if borrowBookResponse.GetSuccess() {
        log.Println("Book borrowed successfully")
    } else {
        log.Println("Failed to borrow book")
    }
}
