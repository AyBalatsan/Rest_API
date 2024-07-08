package main

import (
	"log"

	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/pkg/handler"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
	"github.com/AyBalatsan/Rest_API/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Printf("error occured while running http server: %s", err.Error())
	}
}
