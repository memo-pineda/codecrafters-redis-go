package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
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

	wg := new(sync.WaitGroup)

	for {
		conn, err := listener.Accept()
		fmt.Println("Connection made")

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		wg.Add(1)
		go handleClient(conn, wg)
	}
}

func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		resp := string(buf[:n])

		fmt.Println("Received data", resp)

		if strings.Contains(resp, "PING") {
			conn.Write([]byte("+PONG\r\n"))
		}
	}
}
