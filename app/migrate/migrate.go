package main

import (
	"app/app/database"
	initializers "app/app/ini"
	"app/app/models"
	"fmt"
	"log"
)

func init() {

	config, err := initializers.LoadConfig("../..")
	if err != nil {
		log.Fatal("env not working ", err)

	}

	database.ConnectDB(&config)

}

func main() {
	database.GlobalDB.AutoMigrate(&models.User{})
	database.GlobalDB.AutoMigrate(&models.Posts{})
	fmt.Println("? migration complete")
}
