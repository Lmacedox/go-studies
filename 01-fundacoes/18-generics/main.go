package main

import "fmt"

type x interface {
}

func main() {
	var minhaVar interface{} = "Lucas Matos"

	println(minhaVar.(string))

	// Aqui estamos tentando converter e pegando o resultado.
	// se deu errado o ok é false
	// mas se der certo ok é true
	res, ok := minhaVar.(int)

	fmt.Printf("o valor de res é %v e o resultado de ok é %v\n", res, ok)
}
