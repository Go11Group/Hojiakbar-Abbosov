package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "module/proto"
)

type server struct {
    pb.UnimplementedTranslatorServer
}

var uzbekToEnglish = map[string]string{
    "salom": "hello",
    "dunyo": "world",
}

func (s *server) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
    var englishWords []string
    for _, word := range req.UzbekWords {
        if translation, ok := uzbekToEnglish[word]; ok {
            englishWords = append(englishWords, translation)
        } else {
            englishWords = append(englishWords, word)
        }
    }
    return &pb.TranslateResponse{EnglishWords: englishWords}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterTranslatorServer(s, &server{})
    log.Printf("server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
