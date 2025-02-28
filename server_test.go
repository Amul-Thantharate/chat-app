package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

// Start a test chat server
func startTestServer(t *testing.T) net.Listener {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				t.Logf("Server accept error: %v", err)
				return
			}
			client := newClient(conn, "testUser")
			go handleClient(client)
		}
	}()
	time.Sleep(1 * time.Second) // Allow server to start
	return listener
}

// Test if the server starts successfully
func TestServerStart(t *testing.T) {
	listener := startTestServer(t)
	defer listener.Close()

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	conn.Close()
}

// Test sending and receiving messages
func TestClientMessaging(t *testing.T) {
	listener := startTestServer(t)
	defer listener.Close()

	client1, _ := net.Dial("tcp", "localhost:8080")
	client2, _ := net.Dial("tcp", "localhost:8080")

	defer client1.Close()
	defer client2.Close()

	writer1 := bufio.NewWriter(client1)
	reader2 := bufio.NewReader(client2)

	message := "Hello from client1\n"
	writer1.WriteString(message)
	writer1.Flush()

	received, err := reader2.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to receive message: %v", err)
	}

	if strings.TrimSpace(received) != strings.TrimSpace(message) {
		t.Errorf("Expected %q but got %q", message, received)
	}
}

// Test private messaging
func TestPrivateMessaging(t *testing.T) {
	listener := startTestServer(t)
	defer listener.Close()

	client1, _ := net.Dial("tcp", "localhost:8080")
	client2, _ := net.Dial("tcp", "localhost:8080")

	defer client1.Close()
	defer client2.Close()

	writer1 := bufio.NewWriter(client1)
	reader2 := bufio.NewReader(client2)

	privateMessage := "@testUser Private Hello\n"
	writer1.WriteString(privateMessage)
	writer1.Flush()

	received, err := reader2.ReadString('\n')
	if err != nil {
		t.Fatalf("Failed to receive private message: %v", err)
	}

	if strings.TrimSpace(received) != strings.TrimSpace(privateMessage) {
		t.Errorf("Expected %q but got %q", privateMessage, received)
	}
}

// Test encrypted messaging
func TestEncryptedMessaging(t *testing.T) {
	listener := startTestServer(t)
	defer listener.Close()

	client1, _ := net.Dial("tcp", "localhost:8080")
	client2, _ := net.Dial("tcp", "localhost:8080")

	defer client1.Close()
	defer client2.Close()

	writer1 := bufio.NewWriter(client1)
	reader2 := bufio.NewReader(client2)

	message := "Secret Message"
	encryptedMsg, _ := encryptMessage(message, "key123")

	writer1.WriteString(encryptedMsg + "\n")
	writer1.Flush()

	received, _ := reader2.ReadString('\n')
	decryptedMsg, _ := decryptMessage(strings.TrimSpace(received), "key123")

	if decryptedMsg != message {
		t.Errorf("Expected decrypted message %q but got %q", message, decryptedMsg)
	}
}
