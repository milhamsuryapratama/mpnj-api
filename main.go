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

	kategoriRepo := repositories.NewKategoriRepository(db)
	kategoriUsecase := usecase.NewKategoriUseCase(kategoriRepo)

	api := r.Group("/api")
	handler.KategoriHandlerFunc(api, kategoriUsecase)
	r.Run()
}
