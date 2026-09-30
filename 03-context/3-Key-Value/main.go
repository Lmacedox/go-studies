package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// contexto com valores
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Pegando valor do token do context
	token := ctx.Value("token")
	fmt.Println(token)
}
