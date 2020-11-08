package handler

import (
	"mpnj-api/domain"
	"strconv"

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
	r.POST("/produk", handler.CreateProduk)
	r.GET("/produk/:id", handler.GetProdukByID)
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

// CreateProduk ...
func (p *ProdukHandler) CreateProduk(c *gin.Context) {
	var produk domain.Produk
	var err error

	// binding form to struct
	err = c.Bind(&produk)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	prod, err := p.ProdukUsecase.Create(c, &produk)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})

		c.Abort()
		return
	}

	c.JSON(201, gin.H{
		"message": "sukses",
		"data":    prod,
	})
}

// GetProdukByID ...
func (p *ProdukHandler) GetProdukByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	produk, err := p.ProdukUsecase.GetByID(c, id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{
		"message": "sukses",
		"data":    produk,
	})
}
