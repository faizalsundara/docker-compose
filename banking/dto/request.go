package dto

type UserRequest struct {
	Nama string `json:"nama"`
	NIK  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type SetorSaldo struct {
	NoRekening string `json:"no_rekening"`
	Nominal    int    `json:"nominal"`
}

type TarikSaldo struct {
	NoRekening string `json:"no_rekening"`
	Nominal    int    `json:"nominal"`
}

type DataNorek struct {
	Norek string
}
