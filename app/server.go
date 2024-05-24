package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Listening for connections")

	for {
		conn, err := listener.Accept()
		fmt.Println("Connection made")

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	fmt.Println("Received data", string(buf[:n]))

	commands := strings.Fields(string(buf[:n]))

	for _, command := range commands {
		if strings.Contains(command, "PING") {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}
