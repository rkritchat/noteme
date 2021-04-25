package main

import (
	"log"
	"noteme/internal/config"
	"noteme/internal/repository"
	"noteme/internal/router"
	"noteme/internal/service"
)

func main() {
	cfg := config.NewConfig()
	userRepo := repository.NewUserRepository(cfg.DB)
	userService := service.NewUser(userRepo)
	log.Fatal(router.NewRouter(userService).Listen(":3000"))
}
