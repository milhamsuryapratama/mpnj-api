package usecase

import (
	"context"
	"gorm.io/gorm"
	"mpnj-api/config"
	"mpnj-api/domain"
)

type TransaksiUsecase struct {
	transaksiRepository domain.TransaksiRepository
	db *gorm.DB
}

// NewTransaksiUsecase ...
func NewTransaksiUsecase(t domain.TransaksiRepository) domain.TransaksiUsecase {
	return &TransaksiUsecase{
		transaksiRepository: t,
		db: config.Connect(),
	}
}

// Create ...
func (t *TransaksiUsecase) Create(ctx context.Context, transaksi *domain.Transaksi) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		err := t.transaksiRepository.Create(ctx, transaksi)
		if err != nil {
			return err
		}

		return nil
	})
}

func (t *TransaksiUsecase) Get(ctx context.Context, page int) (domain.TransaksiWithPagination, error) {
	transaksi, err := t.transaksiRepository.Get(ctx, page)
	return transaksi, err
}