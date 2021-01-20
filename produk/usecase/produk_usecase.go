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
func (p *ProdukUsecase) Create(c context.Context, produk *domain.Produk) (err error) {
	err = p.produkRepo.Create(c, produk)
	return
}

// GetByID ...
func (p *ProdukUsecase) GetByID(c context.Context, id int) (prod domain.Produk, err error) {
	prod, err = p.produkRepo.GetByID(c, id)
	return
}

// Update ...
func (p *ProdukUsecase) Update(prod *domain.Produk, id int) (produk domain.Produk, err error) {
	var ctx context.Context
	produk, err = p.produkRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	produk, err = p.produkRepo.Update(prod, id)
	return
}

// Delete ...
func (p *ProdukUsecase) Delete(ctx context.Context, id int) error {
	produk, err := p.produkRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	p.produkRepo.Delete(ctx, &produk, id)
	return nil
}
