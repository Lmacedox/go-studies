package main

import "fmt"

// Da forma atual que a funcao é criada, estamos recebendo os parametros a e b são apenas uma cópia
// Ou seja nao passamos o endereco da memoria aonde ela está guardada, entao mesmo que atualizamos o valor de a dentro de soma, o valor da variavel nao será alterado
func soma(a, b int) int {
	a = 50
	return a + b
}

// Para alterar o valor, devemos receber o endereco de memoria dessas variaveis
func somaWithPointer(a, b *int) int {
	// E para alterar o valor, acessamos o endereco com *
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	fmt.Println(soma(minhaVar1, minhaVar2))
	// Como somaWithPointer precisa de um ponteiro, passamos o endereco da memoria com &
	fmt.Println(somaWithPointer(&minhaVar1, &minhaVar2))
	fmt.Println(minhaVar1)
}
