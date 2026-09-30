package main

/*
Recebendo um genérico, temos o poder de receber diversos parametros diferentes e usa-los conforme quisermos
Abaixo anteriormente recebiamos apenas INT, o que era um problema porque se alguem fosse utilizar essa funcao com float64 iria quebrar
ao invés de duplicar a função apenas utilizamos generics
*/
func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func main() {
	m := map[string]int{
		"Lucas": 10,
		"Diego": 50,
	}

	println(Soma(m))

}
