package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var routes map[string]map[string]func(conn net.Conn)
var body string

func main() {
	body = `
		<html>
			<body>
				%s
			</body>
		</html>
	`

	routes = map[string]map[string]func(conn net.Conn){
		"GET": {
			"/": func(conn net.Conn) {
				b := fmt.Sprintf(body, "<h1>Hello World</h1>")
				response(conn, b)
			},
			"/profile": func(conn net.Conn) {
				b := fmt.Sprintf(body,
					`
						<h1>Perfil</h1>
						<p>Nome: Alexandre</p>
						<p>E-mail: alexandre@gmail.com</p>
					`,
				)
				response(conn, b)
			},
		},
	}

	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	url := strings.Fields(ln)[1]
	fmt.Println("METHOD", method)
	fmt.Println("URL", url)

	route, ok := routes[method]
	if !ok {
		b := fmt.Sprintf(body, "<h1>Página não encontrada</h1>")
		response(conn, b)
		return
	}

	action, ok := route[url]
	if !ok {
		b := fmt.Sprintf(body, "<h1>Página não encontrada</h1>")
		response(conn, b)
		return
	}

	action(conn)
}

func response(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
