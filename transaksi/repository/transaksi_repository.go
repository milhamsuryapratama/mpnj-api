package repository

import (
	"context"
	"gorm.io/gorm"
	"mpnj-api/domain"
)

type TransaksiRepository struct {
	Conn *gorm.DB
}

func NewTransaksiRepository(Conn *gorm.DB) domain.TransaksiRepository {
	return &TransaksiRepository{
		Conn: Conn,
	}
}

func (t *TransaksiRepository) Create(ctx context.Context, transaksi *domain.Transaksi) error {
	t.Conn.Create(&transaksi)
	return nil
}

func (t *TransaksiRepository) Get(ctx context.Context, page int) (domain.TransaksiWithPagination, error) {
	var transaksi domain.TransaksiWithPagination
	offset := (page - 1) * 5
	count := t.Conn.Find(&transaksi.Transaksi)
	t.Conn.Debug().Offset(offset).Limit(5).Preload("User").Preload("TransaksiDetail.Produk").Find(&transaksi.Transaksi)
	transaksi.Total = int(count.RowsAffected)
	transaksi.Page = page
	return transaksi, nil
}