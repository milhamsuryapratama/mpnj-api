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
func (k *KategoriUsecase) Get(c context.Context) ([]domain.Kategori, error) {
	kategori, err := k.kategoriRepo.Get(c)
	return kategori, err
}

// Create ...
func (k *KategoriUsecase) Create(c context.Context, kategori *domain.Kategori) error {
	err := k.kategoriRepo.Create(c, kategori)
	if err != nil {
		return err
	}
	return nil
}

// GetByID ...
func (k *KategoriUsecase) GetByID(c context.Context, id int) (kategori domain.Kategori, err error) {
	kategori, err = k.kategoriRepo.GetByID(c, id)
	return
}

// Update ...
func (k *KategoriUsecase) Update(c context.Context, kat *domain.Kategori, id int) (domain.Kategori, error) {
	_, err := k.kategoriRepo.GetByID(c, id)
	if err != nil {
		return domain.Kategori{}, err
	}

	kategori, err := k.kategoriRepo.Update(c, kat, id)
	return kategori, err
}

// Delete ...
func (k *KategoriUsecase) Delete(c context.Context, id int) (err error) {
	kategori, err := k.kategoriRepo.GetByID(c, id)
	if err != nil {
		return err
	}

	err = k.kategoriRepo.Delete(c, &kategori, id)
	return	err
}
