package main

import (
	"WBproj/cmd/service"
	"WBproj/pkg"
	"WBproj/pkg/logging"
	"WBproj/repository"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("createDB")
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "ec2-34-249-247-7.eu-west-1.compute.amazonaws.com",
		Port:     "5432",
		Username: "kfireyqrkgozaa",
		Password: "31b2140dfdba297c412bda66a9db337c91a8729b17a9791bea82c934ff095d4c",
		DBName:   "d900njt9tj61n8",
		SSLMode:  "require",
	})
	if err != nil {
		logger.Fatal(err)
	}

	repos := repository.NewRepository(db)
	servises := service.NewService(repos)
	handlers := pkg.NewHandler(servises, logger)
	srv := new(pkg.Server)
	cfg := pkg.Config{}
	err = cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		logger.Fatal(err)
	}
	path := cfg.Host + ":" + cfg.Port
	err = srv.Run(path, handlers.InitRoutes())
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("сервер отключен")
}
