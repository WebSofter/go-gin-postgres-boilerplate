package main

import (
	"log"

	"github.com/spf13/viper"
	todo "github.com/websofter/go-gin-postgres-boilerplate"
	"github.com/websofter/go-gin-postgres-boilerplate/pkg/handler"
	"github.com/websofter/go-gin-postgres-boilerplate/pkg/repository"
	"github.com/websofter/go-gin-postgres-boilerplate/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error intializating configs: %s", err.Error())
	}
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
