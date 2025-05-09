package main

import "fmt"

func main() {
	header := []string{"nome", "idade", "email"}
	row := []string{"Alexandre", "30", "alexandre@gmail.com"}

	table := [][]string{header, row}
	fmt.Println(table)

	for _, row := range table {
		for _, column := range row {
			fmt.Printf("%s ", column)
		}
		fmt.Println()
	}
}
