package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func paginaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Pagina principal</h1>")
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public/formulario"))
	mux.Handle("/", http.StripPrefix("/formulario", fs))

	server := &http.Server{ //Servidor personalizado
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Funcionando...")
	log.Fatal(server.ListenAndServe())
}
