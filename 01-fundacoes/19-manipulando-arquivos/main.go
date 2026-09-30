package main

import (
	"fmt"
	"os"
)

func main() {
	// Criando um arquivo chamado "arquivo.txt"
	f, err := os.Create("arquivo.txt")

	if err != nil {
		fmt.Println("Erro ao criar o arquivo")
		panic(err)
	}

	// Escrevendo algo no arquivo
	tamanho, err := f.WriteString("Hello, World!")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes \n", tamanho)

	f.Close()

	// Lendo arquivos
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))
}
