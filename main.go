package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer server.Close()

	fmt.Println("Chat server started on port 8080")
	acceptClients(server)
}
