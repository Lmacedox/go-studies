package main

import "fmt"

// funcoes

func main() {
	total := func() int {
		return sum(10, 10)
	}()

	fmt.Println(total)
}

// funcao variatica
// funcao variatica é uma funcao que nao tem numeros exatos de parametros, podendo receber 1,2, 40, quantos quem chama quiser passar
func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
