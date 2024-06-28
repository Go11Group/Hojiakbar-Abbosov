package main

import (
	"log"
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson47/api_Gateway/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	conn, err := grpc.Dial("localhost: 50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatal(err)
	}
	router := api.Router(*conn)
	panic(router.Run(":50052"))
}