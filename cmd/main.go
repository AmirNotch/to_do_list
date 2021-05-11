package main

import (
	"To_do_list"
	handler "To_do_list/pckg/handler"
	"To_do_list/pckg/repository"
	"To_do_list/pckg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(To_do_list.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occur while running http server: %s", err.Error())
	}
}
