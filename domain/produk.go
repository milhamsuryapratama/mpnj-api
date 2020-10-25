package domain

// Produk ...
type Produk struct {
	IDProduk   int
	NamaProduk string
	Slug       string
	Satuan     string
	Berat      int
	HargaModal int
	HargaJual  int
	Diskon     int
	Stok       int
	Keterangan string
	TipeProduk string
	UserID     int
	Wishlist   int
	Terjual    int
	Kategori   Kategori
	Status     string
}
