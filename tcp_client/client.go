package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := "Hello, TCP server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving:", err)
		return
	}

	receivedMessage := buffer[:n]
	fmt.Printf("Received from server: %s\n", receivedMessage)
}
