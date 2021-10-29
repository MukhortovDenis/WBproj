package main

import (
	"WBproj/pkg"
	"log"
)

func main() {
	handlers := new(pkg.Handler)
	srv := new(pkg.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
