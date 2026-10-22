package main

import (
	"7-packaging/math"
	"fmt"
)

func main() {
	println("Hello, World!")
	// Aqui estamos importando nosso pacote Math
	// Em seguida estamos passando para sua struct os valores que ele pede
	// Em seguida usando seu método
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Add())
}
