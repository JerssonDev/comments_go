package main

import (
	"flag"
	"log"

	"github.com/go-workspace/Comments_Go/migration"
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

}
