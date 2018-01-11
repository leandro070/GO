package main

import (
	"encoding/json"
	"net/http"
)

func FormularioHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/formulario", http.FileServer(http.Dir("public/formulario"))).ServeHTTP(w, r)
	param := r.URL.Query().Get("accion")
	if param == "obtenerProvincias" {
		provincias := getProvincias()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(provincias)
	}

}
