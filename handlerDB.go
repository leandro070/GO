package main

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var settings = mysql.ConnectionURL{
	Host:     "localhost",
	Database: "formulario",
	User:     "root",
	Password: "leandro98237",
}

func GetConnectionDB() sqlbuilder.Database {
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Print("Error al abrir DB: " + err.Error())
	} else {
		log.Print("DB abierta...")
	}
	return sess
}
