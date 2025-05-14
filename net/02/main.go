package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Println(err)
		}

		// goroutine para lidar com mais de uma requisição
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("O código nunca vai chegar aqui porque fica no loop infinito da connection")
}
