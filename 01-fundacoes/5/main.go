package main

import "fmt"

// 	Percorrendo Array

type ID int

var f ID = 1

func main() {
	// Arrays tem posicao fixa
	var meuArray [3]int

	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	// Percorrendo array com for

	for indice, valor := range meuArray {
		fmt.Printf("O valor do indice: %d é %d\n", indice, valor)
	}
}
