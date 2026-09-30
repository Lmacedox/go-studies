package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func main() {
	// Aqui estamos criando a rota
	http.HandleFunc("/", BuscarCEP)

	// Iniciando o servidor HTTP na porta 8080
	http.ListenAndServe(":8080", nil)
}

type Lucas struct {
	Lucas int `json:"lucas"`
}

func BuscarCEP(w http.ResponseWriter, r *http.Request) {
	var lucas Lucas

	req, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(req, &lucas)
	fmt.Println(lucas)
	w.Write([]byte("Hello amigo"))
}
