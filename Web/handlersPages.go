package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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
			Error(w, "Bad Request", http.StatusBadRequest)
			return
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
			Error(w, "Server error", http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)

			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			log.Print("Departamentos servidos")
		}
	}
}

type DNI struct {
	Dni string `json:"dnivalue"`
}
type Response struct {
	Respuesta bool `json:"respuesta"`
}

func PersonaHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("accion")
	if param == "ExistenciaPersona" {
		var dni DNI
		if r.Body == nil {
			Error(w, "Bad Request. Please send a request body", http.StatusBadRequest)
			return
		}

		/*b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}*/ //esto es para usar con Unmarshall, es mas lento que usar Decode

		err := json.NewDecoder(r.Body).Decode(&dni)
		if err != nil {
			Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		cantPers := ExistenciaPersona(dni.Dni)
		respuesta := Response{Respuesta: bool(cantPers)}
		log.Print(respuesta)
		err = json.NewEncoder(w).Encode(respuesta)
		if err != nil {
			Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}

	if param == "GuardarPersona" {

	}
}

func Error(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}
