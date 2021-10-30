package main

import (
	"WBproj/cmd/service"
	"WBproj/pkg"
	"WBproj/repository"
	"log"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "ec2-34-249-247-7.eu-west-1.compute.amazonaws.com",
		Port:     "5432",
		Username: "kfireyqrkgozaa",
		Password: "31b2140dfdba297c412bda66a9db337c91a8729b17a9791bea82c934ff095d4c",
		DBName:   "d900njt9tj61n8",
		SSLMode:  "require",
	})
	if err != nil{
		log.Print(err.Error())
	}
	repos := repository.NewRepository(db)
	servises := service.NewService(repos)
	handlers := pkg.NewHandler(servises)
	srv := new(pkg.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
