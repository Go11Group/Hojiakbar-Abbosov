package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "module/proto"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewTranslatorClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    r, err := c.Translate(ctx, &pb.TranslateRequest{UzbekWords: []string{"salom", "dunyo"}})
    if err != nil {
        log.Fatalf("could not translate: %v", err)
    }
    log.Printf("Translations: %v", r.EnglishWords)
}
