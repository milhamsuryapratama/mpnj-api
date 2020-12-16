package main

import (
	"github.com/gin-gonic/gin"
	"mpnj-api/config"
	"mpnj-api/kategori/delivery/http"
	"mpnj-api/kategori/repository"
	"mpnj-api/kategori/usecase"
)

func main() {
	r := gin.New()
	db := config.Connect()

	kategoriRepo := repository.NewKategoriRepository(db)
	kategoriUsecase := usecase.NewKategoriUseCase(kategoriRepo)

	//produkRepo := repositories.NewProdukRepository(db)
	//produkUsecase := usecase.NewProdukUsecase(produkRepo)

	api := r.Group("/api")
	http.KategoriHandlerFunc(api, kategoriUsecase)
	//handler.ProdukHandlerFunc(api, produkUsecase)

	r.Run()
}
