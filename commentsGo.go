package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/go-workspace/Comments_Go/migration"
	"github.com/go-workspace/Comments_Go/routes"
)

func main() {

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la Base de Datos")
	flag.Parse()

	if migrate == "yes" {

		log.Println("Comenzó la migración...")

		migration.Migrate()

		log.Println("Finalizó la migración!")
	}

	// inician las rutas
	router := routes.InitRoutes()

	//inician los middelwares

	n := negroni.Classic()
	n.UseHandler(router)

	//inicia el servidor
	server := &http.Server{
		Addr:    ":8081",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:8081")

	log.Fatal(server.ListenAndServe())

	log.Println("Fin de la ejecución del programa")

}
