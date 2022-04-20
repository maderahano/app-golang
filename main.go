package main

import (
	"app-golang/config"
	"app-golang/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
