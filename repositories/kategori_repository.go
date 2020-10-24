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

// Create ...
func (c *KategoriRepository) Create(ctx context.Context, k *domain.Kategori) (err error) {
	err = c.Conn.Create(k).Error
	return
}

// GetByID ...
func (c *KategoriRepository) GetByID(ctx context.Context, id int) (domain.Kategori, error) {
	var kategori domain.Kategori
	m := c.Conn.First(&kategori, id).Error
	return kategori, m
}

// Update ...
func (c *KategoriRepository) Update(k *domain.Kategori, id int) (domain.Kategori, error) {
	var kategori domain.Kategori
	kat := c.Conn.First(&kategori, id)
	if kat.Error != nil {
		return kategori, kat.Error
	}

	kategori.NamaKategori = k.NamaKategori
	c.Conn.Save(&kategori)
	return kategori, nil
}
