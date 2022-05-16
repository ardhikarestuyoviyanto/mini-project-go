package repository

import (
	"fmt"
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/model"
	"strings"

	"gorm.io/gorm"
)

type repositoryPegawai struct {
	DB *gorm.DB
}

func (r *repositoryPegawai) GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan {
	var kategoriperizinan []model.APIResponseKategoriPerizinan
	r.DB.Table("kategori_perizinan").Scan(&kategoriperizinan)
	return kategoriperizinan
}

//-------------------------------------------------------------------------------------------------------------------

func (r *repositoryPegawai) CreatePerizinan(perizinan model.Perizinan) bool {
	var kategoriperizinan model.KategoriPerizinan
	r.DB.Table("kategori_perizinan").Select("max_day").Where("id", perizinan.KategoriPerizinanID).Scan(&kategoriperizinan)
	if !lib.ValidationDatePerizinan(perizinan.Start, perizinan.Finish, kategoriperizinan.MaxDay) {
		return false
	}
	r.DB.Create(&perizinan)
	return true
}

func (r *repositoryPegawai) GetAllPerizinan(pegawai_id int) []model.APIResponsePerizinan {
	var perizinan []model.APIResponsePerizinan
	r.DB.Table("users").Select("users.nama AS pegawai_nama", "users.id AS user_id", "kategori_perizinan.id AS kategori_perizinan_id", "kategori_perizinan.name AS kategori_perizinan_nama", "perizinan.catatan", "perizinan.status", "perizinan.start", "perizinan.finish", "perizinan.id").Joins("inner join perizinan on perizinan.user_id = users.id").Joins("inner join kategori_perizinan on perizinan.kategori_perizinan_id=kategori_perizinan.id").Where("perizinan.user_id", pegawai_id).Scan(&perizinan)
	return perizinan
}

func (r *repositoryPegawai) GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan {
	var perizinan model.APIResponsePerizinan
	r.DB.Table("users").Select("users.nama AS pegawai_nama", "users.id AS user_id", "kategori_perizinan.id AS kategori_perizinan_id", "kategori_perizinan.name AS kategori_perizinan_nama", "perizinan.catatan", "perizinan.status", "perizinan.start", "perizinan.finish", "perizinan.id").Joins("inner join perizinan on perizinan.user_id = users.id").Joins("inner join kategori_perizinan on perizinan.kategori_perizinan_id=kategori_perizinan.id").Where("perizinan.id", perizinan_id).Scan(&perizinan)
	return perizinan
}

func (r *repositoryPegawai) UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) bool {
	var kategoriperizinan model.KategoriPerizinan
	r.DB.Table("kategori_perizinan").Select("max_day").Where("id", perizinan.KategoriPerizinanID).Scan(&kategoriperizinan)
	if !lib.ValidationDatePerizinan(perizinan.Start, perizinan.Finish, kategoriperizinan.MaxDay) {
		return false
	}
	r.DB.Model(&model.Perizinan{}).Where("id", perizinan_id).Updates(perizinan)
	return true
}

//----------------------------------------------------------------------------------------------------------------

func (r *repositoryPegawai) ActionInsertAbsenPulang(pegawai_id int, tgl string, pulang string, foto string) {
	rekapAbsen := model.RekapAbsen{
		UserID:     pegawai_id,
		Tanggal:    tgl,
		Pulang:     pulang,
		FotoPulang: foto,
		Keterangan: "Hadir",
	}
	r.DB.Select("UserID", "Tanggal", "Pulang", "FotoPulang", "Keterangan").Create(&rekapAbsen)
}

func (r repositoryPegawai) ActionUpdateAbsenPulang(pegawai_id int, tgl string, pulang string, foto string) {
	rekapAbsen := model.RekapAbsen{
		UserID:     pegawai_id,
		Tanggal:    tgl,
		Pulang:     pulang,
		FotoPulang: foto,
		Keterangan: "Hadir",
	}
	fmt.Println(rekapAbsen)
	r.DB.Model(&model.RekapAbsen{}).Where("user_id", pegawai_id).Where("tanggal", tgl).Updates(rekapAbsen)
}

func (r *repositoryPegawai) ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string) {
	rekapAbsen := model.RekapAbsen{
		UserID:     pegawai_id,
		Tanggal:    tgl,
		Masuk:      masuk,
		FotoMasuk:  foto,
		Keterangan: "Hadir",
	}
	r.DB.Select("UserID", "Tanggal", "Masuk", "FotoMasuk", "Keterangan").Create(&rekapAbsen)
}

func (r *repositoryPegawai) GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string) {
	var rekapabsen model.RekapAbsen
	r.DB.Table("rekap_absen").Select("masuk", "pulang").Where("user_id", user_id).Where("tanggal", tanggal).Scan(&rekapabsen)
	return rekapabsen.Masuk, rekapabsen.Pulang
}

func (r *repositoryPegawai) CheckHariLibur(unitkerja_id int, dayNow string) bool {
	var jamkerja model.JamKerja
	r.DB.Table("unitkerja").Select("jamkerja.harilibur").Joins("inner join jamkerja on jamkerja.id=unitkerja.jamkerja_id").Where("unitkerja.id", unitkerja_id).Scan(&jamkerja)
	arrHariLibur := strings.Split(jamkerja.Harilibur, ",")
	for _, v := range arrHariLibur {
		if v == dayNow {
			return false
		}
	}
	return true
}

func (r *repositoryPegawai) GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string) {
	var jamkerjaDetail model.JamKerjaDetail
	r.DB.Table("unitkerja").Select("jamkerja_detail.mulai_masuk", "jamkerja_detail.batas_masuk", "jamkerja_detail.mulai_pulang", "jamkerja_detail.batas_pulang").Joins("inner join jamkerja on unitkerja.jamkerja_id = jamkerja.id").Joins("inner join jamkerja_detail on jamkerja_detail.jamkerja_id = jamkerja.id").Where("unitkerja.id", unitkerja_id).Where("jamkerja_detail.hari", hari).Scan(&jamkerjaDetail)
	return jamkerjaDetail.MulaiMasuk, jamkerjaDetail.BatasMasuk, jamkerjaDetail.MulaiPulang, jamkerjaDetail.BatasPulang
}

func (r *repositoryPegawai) GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64) {
	var unitkerja model.APIResponseUnitKerja
	r.DB.Table("unitkerja").Select("latidute", "longtidute").Where("id", unitkerja_id).Scan(&unitkerja)
	return unitkerja.Latidute, unitkerja.Longtidute
}

func (r *repositoryPegawai) ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string) {
	rekapabsen := model.RekapAbsen{
		Tanggal:    tgl,
		Masuk:      masuk,
		FotoMasuk:  foto,
		Keterangan: "Hadir",
	}
	r.DB.Model(&model.RekapAbsen{}).Where("user_id", pegawai_id).Where("tanggal", tgl).Updates(rekapabsen)
}

func NewPegawaiRepository(db *gorm.DB) domain.AdapterPegawaiRepository {
	return &repositoryPegawai{
		DB: db,
	}
}
