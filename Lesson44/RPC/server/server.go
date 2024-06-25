package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
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

type UserByIDReq struct {
	Id int
}

var userData []User

type UserServer struct{}

func main() {
	userServer := new(UserServer)

	rpc.Register(userServer)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Listener error: ", err)
	}
	http.Serve(listener, nil)
}

func (t *UserServer) CreateUser(req *UserReq, res *User) error {
	fmt.Println("Users - > ", userData)

	id := len(userData) + 1

	newUser := User{
		Id:   id,
		Name: req.Name,
		Age:  req.Age,
	}

	fmt.Println("User -> ", newUser)

	userData = append(userData, newUser)

	*res = newUser
	return nil
}

func (t *UserServer) GetUserByID(id *int, res *User) error {
	for _, user := range userData {
		if user.Id == *id {
			*res = user
			return nil
		}
	}
	return fmt.Errorf("User with ID %d not found", *id)
}
