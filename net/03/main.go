package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("Connection timeout")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "$ %s", line)
	}
	defer conn.Close()

	fmt.Println("Código chega aqui porque existe o timeout")
}
