package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Iniciamos o contexto
	ctx := context.Background()

	// Utilizando contexto com regra de expiracao
	// E posso cancelalo independente do tempo
	// Esse contexto será cancelado se tiver timeout de 3 segundos, e será cancelado
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel book")
	}
}
