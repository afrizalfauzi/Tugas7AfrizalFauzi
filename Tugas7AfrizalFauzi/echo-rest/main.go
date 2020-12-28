package main

import (
	db "Tugas7AfrizalFauzi/echo-rest/db"
	routes "Tugas7AfrizalFauzi/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
