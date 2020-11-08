package usecase

import (
	"context"
	"mpnj-api/domain"
)

// ProdukUsecase ...
type ProdukUsecase struct {
	produkRepo domain.ProdukRepository
}

// NewProdukUsecase ...
func NewProdukUsecase(p domain.ProdukRepository) domain.ProdukUsecase {
	return &ProdukUsecase{
		produkRepo: p,
	}
}

// Get ...
func (p *ProdukUsecase) Get(c context.Context) (res []domain.Produk, err error) {
	res, err = p.produkRepo.Get(c)
	if err != nil {
		return nil, err
	}

	return
}

// Create ...
func (p *ProdukUsecase) Create(c context.Context, produk *domain.Produk) (prod domain.Produk, err error) {
	prod, err = p.produkRepo.Create(c, produk)
	return
}

// GetByID ...
func (p *ProdukUsecase) GetByID(c context.Context, id int) (prod domain.Produk, err error) {
	prod, err = p.produkRepo.GetByID(c, id)
	return
}
