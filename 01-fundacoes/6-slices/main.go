package main

import "fmt"

// 	Slices

type ID int

var f ID = 1

func main() {
	// Slice tem tamanho dinamico
	// Diferente do array é só não indicar o tamanho do array na declaracao [] com isso voce cria um slice
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Println("--------------------------")

	// Tudo que tiver a direita ela vai sumir, ou seja, eu dropei tudo que tem até a posicao que eu passar, ou seja ele aparecerá vazio
	// Porem a capacidade inicial dele será mantida
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s), s[:0])

	fmt.Println("--------------------------")
	// Tudo que tiver a direita ela vai sumir, ou seja, eu dropei tudo que tem até a posicao que eu passar, ou seja ele aparecerá vazio
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s), s[:4])

	fmt.Println("---------s[2:]-------------")
	// Agora estou fazendo ao contrario, ele irá pegar tudo da esquerda para frente, ou seja, pos posicao 2 pega tudo
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s), s[2:])

	// Adicionando itens no slice
	s = append(s, 100)

	fmt.Println("---------APPEND-------------")
	fmt.Println(s)

}
