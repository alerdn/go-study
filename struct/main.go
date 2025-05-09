package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
	Profissao
}
type Profissao struct {
	Cargo   Cargo
	Salario float64
}

type Cargo int64

func (c Cargo) String() string {
	return [...]string{"Desenvolvedor", "Analista", "Gerente", "Coordenador"}[c]
}

const (
	Desenvolvedor Cargo = iota
	Analista
	Gerente
	Coordenador
)

func main() {
	pessoa := Pessoa{
		Profissao: Profissao{
			Cargo:   Desenvolvedor,
			Salario: 5000.00,
		},
		Nome:  "Maria",
		Idade: 28,
	}

	fmt.Printf("%v\n", pessoa)

}
