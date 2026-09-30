package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) andou() {
	// Aqui estamos alterando o nome porque apontamos o C para um ponteiro
	// Se declararmos assim func (c Cliente) andou() sem o * receberemos apenas uma copia, já com o * (Pointer Reciver) rebemos o endereco de memoria
	// Exatamente igual as funcoes
	c.Nome = "Lucas Matos"
	fmt.Printf("O cliente %s andou", c.Nome)
}

func main() {
	lucas := Cliente{
		Nome: "Lucas",
	}

	lucas.andou()

	fmt.Printf("O valor da struct com nome %v", lucas.Nome)
}
