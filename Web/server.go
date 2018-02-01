package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Print("Listening...")
	log.Fatal(server.ListenAndServe())
}

/*package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db = AbrirConexion()
	buscar("SELECT nombre FROM provincia", "", true)
}
func iniciarServidor() {

}
func buscar(sentencia string, argumentos string, esConsulta bool) {
	defer db.Close()
	if esConsulta && argumentos == "" {
		rows, err := db.Query(sentencia)
		shit(err)
		for rows.Next() {
			var nombre string
			if err := rows.Scan(&nombre); err != nil {
				log.Fatal(err)
			}
			fmt.Println(nombre)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	} else if esConsulta && argumentos != "" {
		rows, err := db.Query(sentencia, argumentos)
		shit(err)
		fmt.Println(rows)
	} else {
		stmt, err := db.Prepare(sentencia)
		shit(err)
		fmt.Println(stmt)

	}

}
func shit(err error) {
	if err != nil {
		log.Print(err)
	}

}
*/
