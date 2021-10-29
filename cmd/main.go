package main

import (
	"WBproj/pkg"
	"log"
)

func main() {
	srv := new(pkg.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatal(err.Error())
	}
}
