package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int
	Saldo  int
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)

	if err != nil {
		panic(err)
	}

	// Abaixo ambos cospem o resultado no terminal
	json.NewEncoder(os.Stdout).Encode(conta)
	fmt.Println(string(res))

	// A regra geral é bem simples:
	// Dado de fora → Código Go: Usa Unmarshal (converte o JSON recebido em variável Go).
	// Código Go → Dado para fora: Usa Marshal
	// (converte a variável Go em JSON para envio).
	// Exemplo de Unmarshal:

	jsonPuro := []byte(`{"Numero":1,"Saldo":100}`)
	var contaX Conta
	// O Unmarshal vai no endereço da variável e atualizar o seu valor com os dados do JSON.
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	fmt.Println(contaX.Saldo)

}
