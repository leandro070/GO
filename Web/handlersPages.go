package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/NotFound", http.FileServer(http.Dir("public/formulario/not_found.html"))).ServeHTTP(w, r)

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
type DatosDelClientePersona struct {
	Nombre               string           `json:"Nombre"`
	Apellido             string           `json:"Apellido"`
	Apellido_Mat         string           `json:"ApellidoMaterno"`
	DNI                  int              `json:"DNI"`
	CUIL                 int              `json:"Cuil"`
	Nacimiento           Fecha_Nacimiento `json:"FechaNacimiento"`
	Telefono_Fijo        int              `json:"Telefono"`
	Telefono_Celular     int              `json:"Celular"`
	Telefono_Alternativo int              `json:"TelefonoAlternativo"`
	Correo               string           `json:"Email"`
	OID_Persona          int
	FK_Genero            int
	EstadoCivil          string
	Direccion            string
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
		log.Print("GuardarPersona")
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Error(w, "Bad Request. Please send a request body", http.StatusBadRequest)
			return
		}

		log.Print(b)

	}
}

func Error(w http.ResponseWriter, err string, code int) {
	http.Error(w, err, code)
}
