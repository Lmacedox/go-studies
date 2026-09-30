package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// padrao pra variavel de contexto é ctx
	ctx := context.Background()

	// se passar um segundo e tiver usando esse contexto, a execucao em si será cancelada
	// nesse caso se o google demorar 1 microsegundo ele vai cancelar a request por conta do timeout
	ctx, cancel := context.WithTimeout(ctx, 1)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		fmt.Println("ERR 1")
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ERR 2")

		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ERR 3")
		panic(err)
	}

	println(string(body))

}
