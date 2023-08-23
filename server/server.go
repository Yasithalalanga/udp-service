package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:40522")
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer conn.Close()
	fmt.Printf("UDP server listening on %s\n", serverAddr)

	buffer := make([]byte, 1024)
	for {
		bytesRead, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		fmt.Printf("Received data from %s: %s\n", addr.String(), string(buffer[:bytesRead]))

		response := []byte("Received your message: " + string(buffer[:bytesRead]))
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println("Error sending response:", err)
		}
	}

}
