package main

import (
	"log"
	"todo-app-go"
	"todo-app-go/pkg/handler"
	"todo-app-go/pkg/repository"
	"todo-app-go/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8001", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
