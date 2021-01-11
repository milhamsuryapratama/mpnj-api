package repository

import (
	"context"
	"mpnj-api/domain"

	"gorm.io/gorm"
)

// ProdukRepository ...
type ProdukRepository struct {
	Conn *gorm.DB
}

// NewProdukRepository ...
func NewProdukRepository(Conn *gorm.DB) domain.ProdukRepository {
	return &ProdukRepository{Conn}
}

// Get ...
func (p *ProdukRepository) Get(ctx context.Context) (res []domain.ProdukWithKategori, err error) {
	var produk []domain.ProdukWithKategori
	p.Conn.Preload("Kategori").Preload("User").Find(&produk)
	return produk, nil
}

// Create ....
func (p *ProdukRepository) Create(ctx context.Context, pr *domain.Produk) (err error) {
	err = p.Conn.Create(pr).Error
	if err != nil {
		return
	}
	return nil
}

// GetByID ...
func (p *ProdukRepository) GetByID(ctx context.Context, id int) (prod domain.Produk, err error) {
	var produk domain.Produk
	err = p.Conn.Preload("User").Preload("Kategori").First(&produk, id).Error
	return produk, err
}

// Update ...
func (p *ProdukRepository) Update(prod *domain.Produk, id int) (domain.Produk, error) {
	var produk domain.Produk
	pro := p.Conn.First(&produk, "id_produk = ?", id)
	if pro.Error != nil {
		return produk, pro.Error
	}

	// p.Conn.Model(&produk).Update("NamaProduk", prod.NamaProduk)
	produk.NamaProduk = prod.NamaProduk

	p.Conn.Save(&produk)
	return *prod, nil
}

func (p *ProdukRepository) Delete(ctx context.Context, produk *domain.Produk, id int) error {
	p.Conn.Delete(&produk, id)
	return nil
}