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

func ProvinciaHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("accion")

	if param == "obtenerProvincia" {
		provincias := getProvincias()
		if data, err := json.Marshal(provincias); err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		} else {
			w.WriteHeader(http.StatusOK)

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			log.Print("Provincias servidas")
		}
	} else if param == "obtenerDepartamento" {
		id := r.URL.Query().Get("id")
		departamentos := getDepartamento(id)
		if data, err := json.Marshal(departamentos); err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		} else {
			w.WriteHeader(http.StatusOK)

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			log.Print("Departamentos servidos")
		}
	}
}

func PersonaHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("accion")
	if param == "ExistenciaPersona" {
		//falta extraer datos del POST
		var dni []byte
		var cantPers []byte

		err := json.NewDecoder(r.Body).Decode(&dni)
		if err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		} else {
			w.WriteHeader(http.StatusOK)

			w.Header().Set("Content-Type", "application/json")
			cantPers = ExistenciaPersona(dni)
			w.Write(cantPers)

		}
	}
}
