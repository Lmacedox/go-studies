package main

import "fmt"

// structs

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estados    string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	// struct com composicao
	// Aqui usamos a struct de Endereco ou seja compondo
	Endereco
	// Ja aqui é criar uma propriedade
	// EnderecoOb Endereco
}

// Aqui criamos um método ligado ao cliente
// Ou seja, dentro do primeiro (), ligamos o c diretamente a struct cliente
// Isso faz com que a funcao desativar seja acessivel através do Cliente.Desativar
func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado!", c.Nome)
}

func main() {
	// Atribuindo uma struct a uma variavel
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 27,
		Ativo: true,
	}

	// Atualizando valor de uma propriedade
	lucas.Nome = "Lucas Matos"

	// Utilizando metodo desativar que foi ligado a struct Cliente
	lucas.Desativar()

	fmt.Printf("Nome %s, Idade %d, Ativo %t", lucas.Nome, lucas.Idade, lucas.Ativo)

}
