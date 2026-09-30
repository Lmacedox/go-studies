package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	// O defer é um statement que atrasa a execucao
	// ou seja, isso será chamado quando a execucao da funcao principal chegar ao final
	// Para um exemplo melhor execute a main abaixo
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}

func main() {
	defer fmt.Println("Primeira linha") // Isso só será executado pós a terceira linha, ou seja por ultimo
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
