package main

import (
	"context"
	"log"
	"time"

	p "grpc/genproto/user/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
 	conn, err := grpc.NewClient(":7070", grpc.WithTransportCredentials(insecure.NewCredentials()))
 	if err != nil {
  		log.Fatalf("did not connect: %v", err)
 	}
 	defer conn.Close()
 	client := p.NewUserServiceClient(conn)

 	newUser := &p.UserReq{
  		Name:  "John Doe",
		Age: 22,
 	}
 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
 	defer cancel()

 	createdUser, err := client.CreateUser(ctx, newUser)
 	if err != nil {
  		log.Fatalf("could not create user: %v", err)
 	}
 	log.Printf("Created User: %v", createdUser)

	userID := &p.Id{Id: createdUser.Id}
 	user, err := client.GetById(ctx, userID)
 	if err != nil {
  		log.Fatalf("could not retrieve user: %v", err)
 	}
 	log.Printf("Retrieved User: %v", user)
}