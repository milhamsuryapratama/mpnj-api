package domain

import (
	"context"
)

// Kategori ...
type Kategori struct {
	IDKategoriProduk int    `gorm:"primaryKey" tag:"id_kategori_produk" json:"id"`
	NamaKategori     string `tag:"nama_kategori" json:"nama_kategori"`
}

// Tabler ...
type Tabler interface {
	TableName() string
}

// TableName ...
func (Kategori) TableName() string {
	return "kategori_produk"
}

// KategoriUsecase ...
type KategoriUsecase interface {
	Get(ctx context.Context) ([]Kategori, error)
	Create(context.Context, *Kategori) error
	GetByID(ctx context.Context, id int) (Kategori, error)
}

// KategoriRepository ...
type KategoriRepository interface {
	Get(ctx context.Context) (res []Kategori, err error)
	Create(ctx context.Context, k *Kategori) error
	GetByID(ctx context.Context, id int) (Kategori, error)
}
