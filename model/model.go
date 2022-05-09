package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int    `gorm:"primaryKey,AUTO_INCREMENT"`
	RoleName string `json:"role_name"`
}

type JamKerja struct {
	gorm.Model
	ID        int    `gorm:"primaryKey,AUTO_INCREMENT"`
	Nama      string `json:"nama"`
	Harilibur string `json:"harilibur"`
}

type JamKerjaDetail struct {
	gorm.Model
	ID          int      `gorm:"primaryKey,AUTO_INCREMENT"`
	Hari        string   `json:"hari"`
	MulaiMasuk  string   `json:"mulai_masuk"`
	BatasMasuk  string   `json:"batas_masuk"`
	MulaiPulang string   `json:"mulai_pulang"`
	BatasPulang string   `json:"batas_pulang"`
	JamkerjaID  int      `json:"jamkerja_id"`
	JamKerja    JamKerja `gorm:"foreignKey:JamkerjaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UnitKerja struct {
	gorm.Model
	ID         int      `gorm:"primaryKey,AUTO_INCREMENT"`
	Nama       string   `json:"nama"`
	Alamat     string   `json:"alamat"`
	Latidute   float64  `json:"latidute"`
	Longtidute float64  `json:"longtidute"`
	JamkerjaID int      `json:"jamkerja_id"`
	JamKerja   JamKerja `gorm:"foreignKey:JamkerjaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Users struct {
	gorm.Model
	ID          int           `gorm:"primaryKey,AUTO_INCREMENT"`
	Nama        string        `json:"nama"`
	Email       string        `json:"email"`
	Password    string        `json:"password"`
	Jabatan     string        `json:"jabatan"`
	Token       string        `json:"token"`
	UnitkerjaID sql.NullInt64 `json:"unitkerja_id"`
	RoleID      int           `json:"role_id"`
	UnitKerja   UnitKerja     `gorm:"foreignKey:UnitkerjaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role        Role          `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type KategoriPerizinan struct {
	gorm.Model
	ID     int    `gorm:"primaryKey,AUTO_INCREMENT"`
	Name   string `json:"name"`
	MaxDay int    `json:"max_day"`
	Type   string `json:"type" gorm:"type:enum('cuti', 'izin')"`
}

type Perizinan struct {
	gorm.Model
	ID                  int               `gorm:"primaryKey,AUTO_INCREMENT"`
	Status              string            `json:"status" gorm:"type:enum('Y','N','P')"`
	Start               string            `json:"start" gorm:"date"`
	Finish              string            `json:"finish" gorm:"date"`
	FilePendukung       string            `json:"file_pendukung"`
	UserID              int               `json:"user_id"`
	KategoriPerizinanID int               `json:"kategori_perizinan_id"`
	Users               Users             `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	KategoriPerizinan   KategoriPerizinan `gorm:"foreignKey:KategoriPerizinanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type RekapAbsen struct {
	gorm.Model
	ID          int            `gorm:"primaryKey,AUTO_INCREMENT"`
	Tanggal     string         `json:"tanggal" gorm:"type:date"`
	Masuk       sql.NullTime   `json:"masuk"`
	Pulang      sql.NullTime   `json:"pulang"`
	FotoMasuk   sql.NullString `json:"foto_masuk"`
	FotoPulang  sql.NullString `json:"foto_pulang"`
	UserID      int            `json:"user_id"`
	PerizinanID sql.NullInt64  `json:"perizinan_id"`
	Users       Users          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Perizinan   Perizinan      `gorm:"foreignKey:PerizinanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "role"
}

func (JamKerja) TableName() string {
	return "jamkerja"
}

func (JamKerjaDetail) TableName() string {
	return "jamkerja_detail"
}

func (UnitKerja) TableName() string {
	return "unitkerja"
}

func (Users) TableName() string {
	return "users"
}

func (KategoriPerizinan) TableName() string {
	return "kategori_perizinan"
}

func (Perizinan) TableName() string {
	return "perizinan"
}

func (RekapAbsen) TableName() string {
	return "rekap_absen"
}

//-------------------------------------
// API RESPONSE MODEL
//-------------------------------------

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
