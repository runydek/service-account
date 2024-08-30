package models

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	Nama       string `json:"nama"`
	NIK        string `json:"nik" gorm:"unique"`
	NoHP       string `json:"no_hp" gorm:"unique"`
	NoRekening string `json:"no_rekening" gorm:"unique"`
	Saldo      int64  `json:"saldo"`
}

type RegisterRequest struct {
	Nama  string `json:"nama" validate:"required"`
	NIK   string `json:"nik" validate:"required"`
	NoHP  string `json:"no_hp" validate:"required"`
}

type TransactionRequest struct {
	NoRekening string `json:"no_rekening" validate:"required"`
	Nominal    int64  `json:"nominal" validate:"required"`
}
