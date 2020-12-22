package main

import (
	"github.com/gin-gonic/gin"
	"mpnj-api/config"
	"mpnj-api/kategori/delivery/http"
	"mpnj-api/kategori/repository"
	"mpnj-api/kategori/usecase"
	produkRepo "mpnj-api/produk/repository"
	produkUsecase "mpnj-api/produk/usecase"
	produkHandler "mpnj-api/produk/delivery/http"
)

func main() {
	r := gin.New()
	db := config.Connect()

	kategoriRepo := repository.NewKategoriRepository(db)
	kategoriUsecase := usecase.NewKategoriUseCase(kategoriRepo)

	produkRepo := produkRepo.NewProdukRepository(db)
	produkUsecase := produkUsecase.NewProdukUsecase(produkRepo)

	api := r.Group("/api")
	http.KategoriHandlerFunc(api, kategoriUsecase)
	produkHandler.ProdukHandlerFunc(api, produkUsecase)

	r.Run()
}
