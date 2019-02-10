package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // solo llama su init
)

// Configuration ...
type configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration ...
func getConfiguration() configuration {

	var c configuration

	file, err := os.Open("./config.json")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)

	if err != nil {
		log.Fatal(err.Error())
	}

	return c

}

// GetConnection ...devuelve la conexion de la base de datos
func GetConnection() *gorm.DB {

	c := getConfiguration()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&l", c.User, c.Password, c.Server, c.Port, c.Database)

	db, err := gorm.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
