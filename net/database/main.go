package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	db := make(map[string]string)
	command := ""
	key := ""
	value := ""

	for scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)

		if len(fields) == 0 {
			fmt.Fprint(conn, "> ")
			continue
		}

		command = fields[0]

		switch command {
		case "GET":
			key = fields[1]
			fmt.Fprintf(conn, "> %s", db[key])
		case "SET":
			key = fields[1]
			value = fields[2]
			db[key] = value
			fmt.Fprint(conn, "> ")
		}
	}
}
