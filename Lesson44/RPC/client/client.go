package main

import (
	"log"
	"net/rpc"
)

type UserReq struct {
	Name string
	Age  int
}

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	user := UserReq{
		Name: "Ahrorbek",
		Age:  27,
	}

	createdUser := User{}

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	err = client.Call("UserServer.CreateUser", user, &createdUser)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}
	log.Println("Created User:", createdUser)

	getUserbyId := User{}

	err = client.Call("UserServer.GetUserByID", 2, &getUserbyId)
	if err != nil {
		log.Fatal("Client invocation error (GetUserByID): ", err)
	}
	log.Println("User by ID:", getUserbyId)
}
