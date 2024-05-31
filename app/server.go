package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	redisparser "github.com/codecrafters-io/redis-starter-go/app/parser"
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

		// Parse the string command into an array
		resp, err := redisparser.ParseBytes(buf, n)

		if err != nil {
			return
		}

		// We get a pointer to an array from ParseBytes, this will
		// just deference it for easier array indexing
		respArray := *resp

		fmt.Printf("parsing result as %v", respArray)

		command := strings.ToUpper(respArray[0])

		if strings.Contains(command, "PING") {
			conn.Write([]byte("+PONG\r\n"))
		} else if strings.Contains(command, "ECHO") {
			if len(respArray) < 2 {
				conn.Write([]byte("Missing 1 argument"))
			} else {
				conn.Write([]byte(respArray[1]))
			}
		}
	}
}
