package usecase

import (
	"context"
	"mpnj-api/domain"
)

// KategoriUsecase ...
type KategoriUsecase struct {
	kategoriRepo domain.KategoriRepository
}

// NewKategoriUseCase ...
func NewKategoriUseCase(k domain.KategoriRepository) domain.KategoriUsecase {
	return &KategoriUsecase{
		kategoriRepo: k,
	}
}

// Get ...
func (k *KategoriUsecase) Get(c context.Context) (res []domain.Kategori, err error) {
	res, err = k.kategoriRepo.Get(c)
	if err != nil {
		return nil, err
	}

	return
}

// Create ...
func (k *KategoriUsecase) Create(c context.Context, kategori *domain.Kategori) (err error) {
	err = k.kategoriRepo.Create(c, kategori)
	return
}

// GetByID ...
func (k *KategoriUsecase) GetByID(c context.Context, id int) (kategori domain.Kategori, err error) {
	kategori, err = k.kategoriRepo.GetByID(c, id)
	return
}

// Update ...
func (k *KategoriUsecase) Update(kat *domain.Kategori, id int) (kategori domain.Kategori, err error) {
	var ctx context.Context
	kategori, err = k.kategoriRepo.GetByID(ctx, id)
	if err != nil {
		return kategori, err
	}

	kategori, err = k.kategoriRepo.Update(kat, id)
	return
}

// Delete ...
func (k *KategoriUsecase) Delete(c context.Context, id int) (err error) {
	err = k.kategoriRepo.Delete(c, id)
	return
}
