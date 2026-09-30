package main

import "fmt"

func main() {
	// Memoria -> Endereco -> Valor

	a := 10
	// Aqui pegamos o valor
	println(a)
	// Aqui pegamos o endereco, ele trará em qual endereco de memoria está sendo guardado
	println(&a)

	// Sempre que usamos * estamos falando do endereco e nao do valor
	var ponteiro *int = &a
	fmt.Println(ponteiro)
	// Atualizando o valor que está guardando neste endereco, ou seja, como ponteiro guarda o endereco de A, o valor de A será atualizado!
	*ponteiro = 20

	// Estamos atribuindo o endereco de A para B
	b := &a
	println(b) // Saida será o valor. Ex: 200x34ad32952118 ->
	// Agora para pegar o valor, nós fazemos uma dereferencia com *
	println(*b) // Saida será o valor. Ex: 20

}
