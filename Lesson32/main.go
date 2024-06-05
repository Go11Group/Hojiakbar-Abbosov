package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error setting up listener:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8080...")

	clients := make(map[net.Conn]bool)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		clients[conn] = true
		go handleConnection(conn, clients) 
	}
}

func handleConnection(conn net.Conn, clients map[net.Conn]bool) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading message:", err)
			delete(clients, conn)
			return
		}

		fmt.Print("Received message: ", string(message))
		for client := range clients {
			if client != conn {
				_, err = client.Write([]byte(message))
				if err != nil {
					fmt.Println("Error writing message:", err)
					delete(clients, client)
				}
			}
		}
	}
}
