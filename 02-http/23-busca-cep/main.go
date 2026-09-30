package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func createFile(endereco ViaCEP) {
	file, err := os.Create("endereco.txt")

	if err != nil {
		panic(err)
	}

	fileFormated, err := json.Marshal(endereco)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(fileFormated))

	defer file.Close()

	if err != nil {
		panic(err)
	}

}

func main() {
	for _, cep := range os.Args[1:] {
		viaCepURL := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
		req, err := http.Get(viaCepURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao bind requisição: %v\n", err)
		}

		fmt.Println(data)
		createFile(data)
		// Gravando endereço no arquivo
	}
}
