package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//* Escrever no servidor
	fmt.Fprintln(conn, "I dial you")

	//* Ler do servidor
	// bs, err := io.ReadAll(conn)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(string(bs))
}
