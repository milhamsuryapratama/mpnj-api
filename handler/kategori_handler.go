package handler

import (
	"mpnj-api/config"
	"mpnj-api/domain"
	"mpnj-api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
	r.POST("/kategori", handler.CreateKategori)
	r.GET("/kategori/:id", handler.GetByID)
	r.PUT("/kategori/:id", handler.UpdateKategori)
	r.DELETE("/kategori/:id", handler.DeleteKategori)
}

// GetKategori ///
func (u *KategoriHandler) GetKategori(c *gin.Context) {
	kategori, err := u.KategoriUsecase.Get(c)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", kategori, 200)
}

// CreateKategori ...
func (u *KategoriHandler) CreateKategori(c *gin.Context) {
	var (
		kategori domain.Kategori
		err error
	)
	err = c.ShouldBindJSON(&kategori)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	err = u.KategoriUsecase.Create(c, &kategori)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", kategori, 201)
}

// GetByID ...
func (u *KategoriHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	kategori, err := u.KategoriUsecase.GetByID(c, id)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", kategori, 200)
}

// UpdateKategori ...
func (u *KategoriHandler) UpdateKategori(c *gin.Context) {
	var (
		kategori domain.Kategori
		err error
	)
	id, _ := strconv.Atoi(c.Param("id"))
	err = c.ShouldBindJSON(&kategori)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	kat, err := u.KategoriUsecase.Update(c, &kategori, id)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", kat, 200)
}

// DeleteKategori ...
func (u *KategoriHandler) DeleteKategori(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := u.KategoriUsecase.Delete(c, id)
	if err != nil {
		utils.Render(c, err.Error(), nil, 400)
		return
	}

	utils.Render(c, "sukses", nil, 200)
}
