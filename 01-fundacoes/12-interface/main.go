package main

import "fmt"

// structs

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estados    string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado!", c.Nome)
}

// O tipo pessoa é uma interface
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (e Empresa) Desativar() {}

func main() {
	lucas := Cliente{
		Nome:  "Lucas",
		Idade: 27,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
	Desativacao(&lucas)

}
