package main

import "fmt"

func main() {
	literalObject := map[string]string{
		"nome": "Alexandre",
	}

	fmt.Println(literalObject)

	//* Não funciona porque o map não aponta para nada
	// var nilObject map[string]string
	// nilObject["noem"] = "Alexandre"
	// fmt.Println(nilObject)

	//* Agora sim
	initializedObject := make(map[string]string)
	initializedObject["nome"] = "Alexandre"

	if email, ok := initializedObject["email"]; ok {
		fmt.Printf("E-mail: %s", email)
	} else {
		fmt.Println("E-mail não existe. Adicionando...")
		initializedObject["email"] = "alexandre@mail.com"
		fmt.Println(initializedObject["email"])
		delete(initializedObject, "email")
		fmt.Println("E-mail removido.")
		fmt.Println(initializedObject)
	}
}
