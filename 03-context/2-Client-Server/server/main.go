package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Pegando context que veio da requisicao
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(5 * time.Second):
		// O log é impresso no stdout (Terminal)
		log.Println("Request processada com sucesso")
	case <-ctx.Done():
		log.Println("Request cancelada pelo cliente")
		http.Error(w, "Request cancelada pelo cliente", http.StatusRequestTimeout)

	}
	w.Write([]byte("Hello!"))
}
