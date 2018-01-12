package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

func FormularioHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/formulario", http.FileServer(http.Dir("public/formulario"))).ServeHTTP(w, r)

}
func GetProvinciaHandler(w http.ResponseWriter, r *http.Request) {
	provincias := getProvincias()
	if data, err := json.Marshal(provincias); err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	} else {
		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
		log.Print("Provincias servidas")

	}
}
