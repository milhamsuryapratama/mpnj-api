package domain

import (
	"context"
	"time"
)

type Transaksi struct {
	IDTransaksi      int               `gorm:"primaryKey" tag:"id_transaksi" json:"id"`
	KodeTransaksi    string            `tag:"kode_transaksi" form:"kode_transaksi" json:"kode_transaksi"`
	UserID           int               `sql:"index" tag:"user_id" form:"user_id" json:"user_id"`
	User             User              `gorm:"foreignKey:user_id" tag:"user" json:"pembeli"`
	WaktuTransaksi   time.Time         `json:"waktu_transaksi"`
	TotalBayar       int               `tag:"total_bayar" form:"total_bayar" json:"total_bayar"`
	StatusTransaksi  string            `tag:"status_transaksi" form:"status_transaksi" json:"status_transaksi"`
	BatasTransaksi   time.Time         `tag:"batas_transaksi" form:"batas_transaksi" json:"batas_transaksi"`
	ProsesPembayaran string            `tag:"proses_pembayaran" form:"proses_pembayaran" json:"proses_pembayaran"`
	To               string            `tag:"to" form:"to" json:"to"`
	TransaksiDetail  []TransaksiDetail `gorm:"foreignKey:TransaksiID"`
}

type TransaksiWithPagination struct {
	Transaksi []Transaksi `tag:"transaksi" json:"transaksi"`
	Total     int         `tag:"total" json:"total"`
	Page      int         `tag:"page" json:"page"`
}

type TransaksiUsecase interface {
	Create(ctx context.Context, transaksi *Transaksi) error
	Get(ctx context.Context, offset int) (TransaksiWithPagination, error)
}

type TransaksiRepository interface {
	Create(ctx context.Context, transaksi *Transaksi) error
	Get(ctx context.Context, page int) (TransaksiWithPagination, error)
}
