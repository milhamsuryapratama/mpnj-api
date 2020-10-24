package main

import (
	"mpnj-api/config"
	"mpnj-api/handler"
	"mpnj-api/repositories"
	"mpnj-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	db := config.Connect()

	repo := repositories.NewKategoriRepository(db)
	usecase := usecase.NewKategoriUseCase(repo)

	api := r.Group("/api")
	handler.KategoriHandlerFunc(api, usecase)
	r.Run()
}
