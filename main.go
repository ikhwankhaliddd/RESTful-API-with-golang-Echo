package main

import (
	"github.com/ikhwankhaleed/echo-rest/db"
	"github.com/ikhwankhaleed/echo-rest/routes"
)

func main() {

	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
