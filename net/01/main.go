package main

import (
	"fmt"
	"io"
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

		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%s", "Well, I hope!")

		conn.Close()
	}
}
