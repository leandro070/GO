package main

import "log"

type Persona struct {
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
	FK_EstadoCivil       int
	FK_Domicilio         int
}
type Fecha_Nacimiento struct {
	Anio string
	Mes  string
	Dia  string
}

func ExistenciaPersona(dni string) bool {
	db := GetConnectionDB() //traigo la conexion
	defer db.Close()
	cantPers, err := db.
		Collection("persona").
		Find("dni", dni).
		Count()
	log.Print("Cantidad de personas:", cantPers)
	if err != nil {
		log.Print("Error al confirmar existencia de persona: ", err)
	}
	if cantPers > 0 {
		return true
	} else {
		return false
	}

}
