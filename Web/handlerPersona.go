package main

import "log"

type Persona struct {
	Nombre               string
	Apellido             string
	Apellido_Mat         string
	DNI                  int `json:"DNI"`
	CUIL                 int
	Nacimiento           Fecha_Nacimiento
	Telefono_Fijo        int
	Telefono_Celular     int
	Telefono_Alternativo int
	Correo               string
	OID_Persona          int
	FK_Genero            int
	FK_EstadoCivil       int
	FK_Domicilio         int
}
type Fecha_Nacimiento struct {
	Anio string
	Mes  string
	Dia  string
}

func ExistenciaPersona(dni []byte) []byte {
	db := GetConnectionDB() //traigo la conexion
	defer db.Close()
	cantPers, err := db.
		Collection("persona").
		Find("dni", dni).
		Count()
	if err != nil {
		log.Print("Error al confirmar existencia de persona: ", err)
	}
	return make([]byte, cantPers)
}
