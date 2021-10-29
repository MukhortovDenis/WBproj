package main

import (
	"WBproj/cmd/service"
	"WBproj/pkg"
	"WBproj/repository"
	"log"
)

func main() {
	repos := repository.NewRepository()
	servises := service.NewService(repos)
	handlers := pkg.NewHandler(servises)
	srv := new(pkg.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
