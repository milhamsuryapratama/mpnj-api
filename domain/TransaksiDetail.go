package domain

type TransaksiDetail struct {
	IDTransaksiDetail int       `gorm:"primaryKey" tag:"id_transaksi_detail" json:"id"`
	TransaksiID       int       `sql:"index" tag:"transaksi_id" form:"transaksi_id" json:"transaksi_id"`
	Transaksi         Transaksi `gorm:"foreignKey:transaksi_id" tag:"transaksi" json:"-"`
	ProdukID          int       `sql:"index" tag:"produk_id" form:"produk_id" json:"produk_id"`
	Produk            Produk    `gorm:"foreignKey:produk_id" tag:"produk" json:"produk"`
	Jumlah            int       `tag:"jumlah" form:"jumlah" json:"jumlah"`
	Diskon            int       `tag:"diskon" form:"diskon" json:"diskon"`
	Subtotal          int       `tag:"subtotal" form:"subtotal" json:"subtotal"`
	StatusOrder       string    `tag:"status_order" form:"status_order" json:"status_order"`
	UserID            int       `sql:"index" tag:"user_id" form:"user_id" json:"user_id"`
	User              User      `gorm:"foreignKey:user_id" tag:"user" json:"user"`
}
