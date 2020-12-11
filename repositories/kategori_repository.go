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
func (c *KategoriRepository) Get(ctx context.Context) ([]domain.Kategori, error) {
	var kategori []domain.Kategori
	c.Conn.Find(&kategori)
	return kategori, nil
}

// Create ...
func (c *KategoriRepository) Create(ctx context.Context, kategori *domain.Kategori) error {
	err := c.Conn.Create(kategori).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByID ...
func (c *KategoriRepository) GetByID(ctx context.Context, id int) (domain.Kategori, error) {
	var kategori domain.Kategori
	m := c.Conn.First(&kategori, id).Error
	return kategori, m
}

// Update ...
func (c *KategoriRepository) Update(ctx context.Context, k *domain.Kategori, id int) (domain.Kategori, error) {
	var kategori domain.Kategori
	err := c.Conn.First(&kategori, id).Error
	if err != nil {
		return domain.Kategori{}, err
	}

	kategori.NamaKategori = k.NamaKategori
	c.Conn.Save(&kategori)
	return kategori, nil
}

// Delete ...
func (c *KategoriRepository) Delete(ctx context.Context, id int) error {
	var kategori domain.Kategori
	err := c.Conn.First(&kategori, id).Error
	if err != nil {
		return err
	}

	c.Conn.Delete(&kategori, id)
	return err
}
