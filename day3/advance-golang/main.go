package main

import (
	"advance-golang/config"
	"advance-golang/routes"
)

func init() {
	config.InitDB()
}

func main() {
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
