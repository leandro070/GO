package main

import (
	"log"
)

type Departamento struct {
	Nombre          string `db:"nombre" json:"nombre"`
	ID_Departamento string `db:"OID_departamento" json:"id_departamento"`
}

var departamentos []Departamento

func obtenerDepartamentosDB(id_prov string) {
	db := GetConnectionDB() //traigo la conexion
	defer db.Close()
	dptoSelector := db.
		Select("nombre", "OID_departamento").
		From("departamento").
		Where("FK_provincia", id_prov)
	err := dptoSelector.All(&departamentos)
	if err != nil {
		log.Print("Error al recuperar Departamento:", err)
	}
}

func getDepartamento(id_prov string) []Departamento {
	obtenerDepartamentosDB(id_prov)
	return departamentos
}
func printDepartamentos() {
	log.Print(departamentos)
}
