package redisparser

import (
	"errors"
	"fmt"
	"strings"
)

func ParseBytes(bytes []byte, size int) (*[]string, error) {
	resp := string(bytes[:size])

	fmt.Println("Received data", resp)

	parts := strings.Split(resp, `\r\n`)

	arrayDetails := parts[0]

	if arrayDetails[0] == '*' {
		// We don't care for the size atm, but can be an optimization we can make
		// For now, we expect a "*", indicating an array of commands
		command := []string{} // this will store list of commands + args

		remaining := parts[1:]

		idx := 0

		for idx < len(remaining)-1 {
			// Append every command + args to the array
			content := remaining[idx+1]

			command = append(command, content)
			idx += 2
		}

		return &command, nil
	}

	return nil, errors.New("can't parse bytes")
}
