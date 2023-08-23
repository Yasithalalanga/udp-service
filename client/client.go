package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:40522")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	message := []byte("Hello, UDP Server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data:", err)
		os.Exit(1)
	}

	buffer := make([]byte, 1024)
	bytesRead, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		os.Exit(1)
	}

	fmt.Println("Response from server:", string(buffer[:bytesRead]))
}
