package repository

import (
	"mini-project-go/domain"
	"mini-project-go/model"

	"gorm.io/gorm"
)

type repositoryReportPegawai struct {
	DB *gorm.DB
}

func (r *repositoryReportPegawai) GetBulanan(user_id int, month int) []model.APIResponseRekapAbsen {
	var rekapabsen []model.APIResponseRekapAbsen
	r.DB.Table("users").Select("rekap_absen.id", "rekap_absen.masuk", "rekap_absen.pulang", "rekap_absen.tanggal", "rekap_absen.foto_masuk", "rekap_absen.foto_pulang", "rekap_absen.keterangan", "users.nama AS pegawai_name").Joins("inner join rekap_absen on rekap_absen.user_id = users.id").Where("users.id", user_id).Where("MONTH(tanggal) = ?", month).Order("rekap_absen.tanggal desc").Scan(&rekapabsen)
	return rekapabsen
}

func (r *repositoryReportPegawai) GetRangeDay(user_id int, start string, finish string) []model.APIResponseRekapAbsen {
	var rekapabsen []model.APIResponseRekapAbsen
	r.DB.Table("users").Select("rekap_absen.id", "rekap_absen.masuk", "rekap_absen.pulang", "rekap_absen.tanggal", "rekap_absen.foto_masuk", "rekap_absen.foto_pulang", "rekap_absen.keterangan", "users.nama AS pegawai_name").Joins("inner join rekap_absen on rekap_absen.user_id = users.id").Where("users.id", user_id).Where("tanggal BETWEEN ? AND ?", start, finish).Order("rekap_absen.tanggal asc").Scan(&rekapabsen)
	return rekapabsen
}

func NewReportPegawaiRepository(db *gorm.DB) domain.AdapterPegawaiReportRepository {
	return &repositoryReportPegawai{
		DB: db,
	}
}
