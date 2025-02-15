package model

type UserData struct {
	Nama string `json:"nama"`
	NIK  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type Saldo struct {
	Saldo      int    `json:"saldo"`
	NoRekening string `json:"no_rekening"`
}
