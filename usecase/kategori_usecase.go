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
