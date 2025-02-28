package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

// Client struct represents a chat client
type Client struct {
	conn net.Conn
	name string
}

var clients = make(map[string]*Client)
var mutex sync.Mutex

// Create a new client instance
func newClient(conn net.Conn, name string) *Client {
	return &Client{conn: conn, name: name}
}

// Handle client communication
func handleClient(client *Client) {
	defer func() {
		mutex.Lock()
		delete(clients, client.name)
		mutex.Unlock()
		client.conn.Close()
		fmt.Println(client.name, "disconnected")
	}()

	for {
		buffer := make([]byte, 1024)
		n, err := client.conn.Read(buffer)
		if err != nil {
			return
		}

		message := strings.TrimSpace(string(buffer[:n]))

		// Private messaging: format "@username message"
		if strings.HasPrefix(message, "@") {
			parts := strings.SplitN(message, " ", 2)
			if len(parts) < 2 {
				client.conn.Write([]byte("Invalid private message format. Use: @username message\n"))
				continue
			}

			targetName := parts[0][1:] // Extract username
			privateMsg := parts[1]

			mutex.Lock()
			targetClient, exists := clients[targetName]
			mutex.Unlock()

			if exists {
				targetClient.conn.Write([]byte(fmt.Sprintf("[Private] %s: %s\n", client.name, privateMsg)))
				client.conn.Write([]byte(fmt.Sprintf("[Private to %s]: %s\n", targetName, privateMsg)))
			} else {
				client.conn.Write([]byte("User not found!\n"))
			}
			continue
		}

		// Broadcast message to all clients
		broadcastMessage(fmt.Sprintf("%s: %s\n", client.name, message), client.name)
	}
}

// Broadcast message to all clients
func broadcastMessage(message, sender string) {
	mutex.Lock()
	defer mutex.Unlock()
	for name, client := range clients {
		if name != sender {
			client.conn.Write([]byte(message))
		}
	}
}

// Accept new clients
func acceptClients(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		conn.Write([]byte("Enter your username: "))
		buffer := make([]byte, 1024)
		n, _ := conn.Read(buffer)
		username := strings.TrimSpace(string(buffer[:n]))

		mutex.Lock()
		clients[username] = newClient(conn, username)
		mutex.Unlock()

		fmt.Println(username, "joined the chat")
		go handleClient(clients[username])
	}
}
