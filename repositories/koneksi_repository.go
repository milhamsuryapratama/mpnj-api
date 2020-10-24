package repositories

import (
	"context"
	"mpnj-api/domain"

	"gorm.io/gorm"
)

// KategoriRepository ...
type KategoriRepository struct {
	Conn *gorm.DB
}

// NewKategoriRepository ...
func NewKategoriRepository(Conn *gorm.DB) domain.KategoriRepository {
	return &KategoriRepository{Conn}
}

// Get ...
func (c *KategoriRepository) Get(ctx context.Context) (res []domain.Kategori, err error) {
	var kategori []domain.Kategori
	c.Conn.Find(&kategori)
	return kategori, nil
}
