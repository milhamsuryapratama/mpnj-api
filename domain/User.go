package domain

// User ...
type User struct {
	IDUser      int    `gorm:"primaryKey" tag:"id_user" json:"id"`
	NamaLengkap string `tag:"nama_lengkap" form:"nama_lengkap" json:"nama_lengkap"`
	Username    string `tag:"username" form:"username" json:"username"`
	Password    string `gorm:"-" tag:"password" form:"password" json:"password"`
	NomorHp     string `tag:"nomor_hp" form:"nomor_hp" json:"nomor_hp"`
	Email       string `tag:"email" form:"email" json:"email"`
}
