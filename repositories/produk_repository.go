package repositories
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
func (p *ProdukRepository) Get(ctx context.Context) (res []domain.Produk, err error) {
	var produk []domain.Produk
	p.Conn.Preload("Kategori").Preload("User").Find(&produk)
	return produk, nil
}
