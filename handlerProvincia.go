package main

import (
	"log"
)

type Provincia struct {
	Nombre string `db:"nombre"`
	ID     string `db:"OID_provincia"`
}

var provincias []Provincia

func obtenerProviciasDB() {
	db := GetConnectionDB() //tarigo la conexion
	defer db.Close()
	provinciaCollection := db.Collection("provincia") //selecciono la tabla
	res := provinciaCollection.Find()                 //hago un SELECT * FROM provincia
	err := res.All(&provincias)                       //guardo todo en provincias
	if err != nil {
		log.Fatalf("res.All(): %q\n", err)
	}
}

func getProvincias() []Provincia {
	if provincias == nil {
		obtenerProviciasDB()
	}
	return provincias
}
func printProvincias() {
	log.Print(provincias)
}
