package main

import (
	"app/app/database"
	initializers "app/app/ini"
	"app/app/router"
	"log"
)

func init() {

	config, err := initializers.LoadConfig("..")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)

}

func main() {
	r := router.SetUpRouter()

	r.Run(":8080")
}
