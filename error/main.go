package main

import (
	"log"
	"study/error/customexception"
)

func login() error {

	return customexception.New(401, "Senha incorreta")
}

func main() {
	if err := login(); err != nil {
		log.Fatalln(err.Error())
	}
}
