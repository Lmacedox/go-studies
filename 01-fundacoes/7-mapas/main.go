package main

import "fmt"

// 	Maps

func main() {
	// Criando Map manual
	salarios := map[string]int{
		"Wesley": 1500,
		"Lucas":  2500,
	}

	// Criando Map com make
	salWithMake := make(map[string]int)

	fmt.Println(salWithMake)

	fmt.Println(salarios)
	fmt.Println(salarios["Lucas"])

	// Adicionando novo item
	salarios["Diogo"] = 120

	// Deletando
	delete(salarios, "Wesley")
}
