package handler

import (
	"mpnj-api/config"
	"mpnj-api/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// KategoriHandler ...
type KategoriHandler struct {
	KategoriUsecase domain.KategoriUsecase
	Conn            *gorm.DB
}

// KategoriHandlerFunc ...
func KategoriHandlerFunc(r *gin.RouterGroup, kategori domain.KategoriUsecase) {
	handler := &KategoriHandler{
		KategoriUsecase: kategori,
		Conn:            config.Connect(),
	}

	r.GET("/kategori", handler.GetKategori)
}

// GetKategori ///
func (u *KategoriHandler) GetKategori(c *gin.Context) {
	listKategori, err := u.KategoriUsecase.Get(c)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"message": "sukses",
		"data":    listKategori,
	})
}
