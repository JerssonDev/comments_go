package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/go-workspace/Comments_Go/commons"
	"github.com/go-workspace/Comments_Go/migration"
	"github.com/go-workspace/Comments_Go/routes"
)

func main() {

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la Base de Datos")
	flag.IntVar(&commons.Port, "port", 80, "Puerto para el servidor")
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
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado el servidor en http://localhost:%d\n", commons.Port)

	log.Fatal(server.ListenAndServe())

	log.Println("Fin de la ejecución del programa")

}
