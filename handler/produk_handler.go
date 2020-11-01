package handler

import (
	"mpnj-api/domain"

	"github.com/gin-gonic/gin"
)

// ProdukHandler ...
type ProdukHandler struct {
	ProdukUsecase domain.ProdukUsecase
}

// ProdukHandlerFunc ...
func ProdukHandlerFunc(r *gin.RouterGroup, produk domain.ProdukUsecase) {
	handler := &ProdukHandler{
		ProdukUsecase: produk,
	}

	r.GET("/produk", handler.GetProduk)
}

// GetProduk ...
func (p *ProdukHandler) GetProduk(c *gin.Context) {
	listProduk, err := p.ProdukUsecase.Get(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "sukses",
		"data":    listProduk,
	})
}
