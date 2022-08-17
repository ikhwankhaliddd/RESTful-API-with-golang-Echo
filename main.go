package main

import (
	"os"

	"github.com/ikhwankhaleed/echo-rest/db"
	"github.com/ikhwankhaleed/echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
