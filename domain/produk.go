package domain

import "context"

// Produk ...
type Produk struct {
	IDProduk         int      `gorm:"primaryKey" tag:"id_produk" json:"id"`
	NamaProduk       string   `tag:"nama_produk" form:"nama_produk" json:"nama_produk"`
	Slug             string   `tag:"slug" form:"slug" json:"slug"`
	Satuan           string   `tag:"satuan" form:"satuan" json:"satuan"`
	Berat            int      `tag:"berat" form:"berat" json:"berat"`
	HargaModal       int      `tag:"harga_modal" form:"harga_modal" json:"harga_modal"`
	HargaJual        int      `tag:"harga_jual" form:"harga_jual" json:"harga_jual"`
	Diskon           int      `tag:"diskon" form:"diskon" json:"diskon"`
	Stok             int      `tag:"stok" form:"stock" json:"stok"`
	Keterangan       string   `tag:"keterangan" form:"keterangan" json:"keterangan"`
	TipeProduk       string   `tag:"tipe_produk" form:"tipe_produk" json:"tipe_produk"`
	UserID           int      `sql:"index" tag:"user_id" form:"user_id" json:"user_id"`
	Wishlist         int      `tag:"wishlist" form:"wishlist" json:"wishlist"`
	Terjual          int      `tag:"terjual" form:"terjual" json:"terjual"`
	KategoriProdukID int      `sql:"index" form:"kategori_produk_id" tag:"kategori_produk_id" json:"kategori_produk_id"`
	Status           string   `tag:"status" form:"status" json:"status"`
}

type ProdukWithKategori struct {
	Produk
	Kategori         Kategori `gorm:"foreignKey:kategori_produk_id" tag:"kategori" json:"kategori"`
	User             User     `gorm:"foreignKey:user_id" tag:"user" json:"penjual"`
}

// ProdukUsecase ...
type ProdukUsecase interface {
	Get(ctx context.Context) ([]ProdukWithKategori, error)
	Create(context.Context, *Produk) error
	GetByID(ctx context.Context, id int) (Produk, error)
	Update(produk *Produk, id int) (Produk, error)
	Delete(ctx context.Context, id int) error
}

// ProdukRepository ...
type ProdukRepository interface {
	Get(ctx context.Context) (res []ProdukWithKategori, err error)
	Create(ctx context.Context, p *Produk) error
	GetByID(ctx context.Context, id int) (Produk, error)
	Update(produk *Produk, id int) (Produk, error)
	Delete(ctx context.Context, produk *Produk, id int) error
}
