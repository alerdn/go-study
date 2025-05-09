package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Pessoa struct {
	Nome  string
	Email string
	Idade int
}

func (p *Pessoa) SaveData(w io.Writer) {
	fmt.Fprintf(w, "%s;%s;%d\n", p.Nome, p.Email, p.Idade)
}

func gerarHeader(p Pessoa) string {
	pessoaJson, _ := json.Marshal(p)
	pessoaMap := make(map[string]any)
	json.Unmarshal(pessoaJson, &pessoaMap)

	header := ""
	for key := range pessoaMap {
		header += key + ";"
	}
	header = header[:len(header)-1] + "\n"

	return header
}

func main() {
	pessoa := Pessoa{
		Nome:  "Alexandre",
		Email: "ale@mail.com",
		Idade: 25,
	}

	header := gerarHeader(pessoa)

	// Save data to buffer
	buffer := bytes.NewBufferString(header)
	pessoa.SaveData(buffer)
	fmt.Println(buffer)

	// Save data to file
	file, err := os.Create("pessoas.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(header)
	pessoa.SaveData(file)
}
