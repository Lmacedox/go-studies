package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	// aqui definimos o tempo de espera
	c := http.Client{Timeout: 1}
	resp, err := c.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
