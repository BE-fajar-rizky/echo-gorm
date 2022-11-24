package main

import (
	"fajar/apiMvc/config"
	"fajar/apiMvc/routes"
)

func main() {
	config.InitDB()
	config.InitialMigration()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))

}
