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
