package main

import (
	"github.com/rogudator/tuvebes-backend/internal/entity"
	"github.com/rogudator/tuvebes-backend/internal/repository"
	"github.com/rogudator/tuvebes-backend/internal/transport"
)

func main() {
	r := repository.NewRepository()
	r.CreateTuvebe(entity.Tuvebe{
		Date:     "2024.01.07",
		TimeFrom: "21:00",
		TimeTo:   "01:03",
		Title:    "Coding tuvebe mvp",
	})
	handler := transport.NewTransport(r)
	router := handler.InitRoutes()
	
	router.Run("0.0.0.0:8081")
}