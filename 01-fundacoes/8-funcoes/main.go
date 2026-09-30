package main

import "fmt"

// funcoes

func main() {
	valSoma := sum(10, 10)
	fmt.Println(valSoma)

	valGreater, greater := sumGreaterThanFifty(30, 20)

	if greater == true {
		fmt.Printf("Parabéns!!, o valor %d é maior ou igual á 50", valGreater)
	} else {
		fmt.Printf("Infelizmente o valor %d é menor que 50", valGreater)
	}

}

// funcao com devolucao de int
func sum(a, b int) int {
	return a + b
}

// funcao com devolucao de dois valores
func sumGreaterThanFifty(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}
