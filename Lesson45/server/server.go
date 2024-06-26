package main

import (
    "context"
    "database/sql"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson45/protos"
    _ "github.com/lib/pq"
    "github.com/google/uuid"
)

const (
    port     = ":50051"
    host     = "localhost"
    user     = "postgres"
    password = "root"
    dbname   = "go11"
)

type server struct {
    pb.UnimplementedLibraryServiceServer
    db *sql.DB
}

func (s *server) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
    bookID := uuid.New().String()
    _, err := s.db.Exec("INSERT INTO books (book_id, title, author, year_published) VALUES ($1, $2, $3, $4)",
        bookID, req.GetTitle(), req.GetAuthor(), req.GetYearPublished())
    if err != nil {
        return nil, err
    }
    return &pb.AddBookResponse{BookId: bookID}, nil
}

func (s *server) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
    rows, err := s.db.Query("SELECT book_id, title, author, year_published FROM books WHERE title ILIKE $1 OR author ILIKE $2",
        "%"+req.GetQuery()+"%", "%"+req.GetQuery()+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []*pb.Book
    for rows.Next() {
        var book pb.Book
        if err := rows.Scan(&book.BookId, &book.Title, &book.Author, &book.YearPublished); err != nil {
            return nil, err
        }
        books = append(books, &book)
    }
    return &pb.SearchBookResponse{Books: books}, nil
}

func (s *server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
    _, err := s.db.Exec("INSERT INTO borrows (book_id, user_id) VALUES ($1, $2)", req.GetBookId(), req.GetUserId())
    if err != nil {
        return &pb.BorrowBookResponse{Success: false}, err
    }
    return &pb.BorrowBookResponse{Success: true}, nil
}

func main() {
    connStr := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterLibraryServiceServer(s, &server{db: db})

    log.Printf("Server is running on port %s", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

