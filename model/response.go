package model

type APIResponseJamKerja struct {
	ID        int    `json:"ID"`
	Nama      string `json:"nama"`
	Harilibur string `json:"harilibur"`
}

type APIResponseJamKerjaDetail struct {
	Hari        string `json:"hari"`
	MulaiMasuk  string `json:"mulai_masuk"`
	BatasMasuk  string `json:"batas_masuk"`
	MulaiPulang string `json:"mulai_pulang"`
	BatasPulang string `json:"batas_pulang"`
}

type APIResponseUnitKerja struct {
	ID           int     `json:"ID"`
	Nama         string  `json:"nama"`
	Alamat       string  `json:"alamat"`
	Latidute     float64 `json:"latidute"`
	Longtidute   float64 `json:"longtidute"`
	JamkerjaID   int     `json:"jamkerja_id"`
	JamkerjaNama string  `json:"jamkerja_nama"`
}

type APIResponsePegawai struct {
	ID            int    `json:"ID"`
	Nama          string `json:"nama"`
	Email         string `json:"email"`
	Jabatan       string `json:"jabatan"`
	UnitkerjaId   int    `json:"unitkerja_id"`
	UnitkerjaNama string `json:"unitkerja_nama"`
}

type APIResponseKategoriPerizinan struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	MaxDay int    `json:"max_day"`
	Type   string `json:"type"`
}
