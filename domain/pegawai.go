package domain

import "mini-project-go/model"

type AdapterPegawaiRepository interface {
	ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionUpdateAbsenPulang(pegawai_id int, tgl string, pulang string, foto string)
	ActionInsertAbsenPulang(pegawai_id int, tgl string, pulang string, foto string)
	GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64)
	GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string)
	CheckHariLibur(unitkerja_id int, dayNow string) bool
	GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string)
	//-------------------------------------------------------------------------------------------------------
	CreatePerizinan(perizinan model.Perizinan) bool
	GetAllPerizinan(pegawai_id int) []model.APIResponsePerizinan
	GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan
	UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) bool
	//-------------------------------------------------------------------------------------------------------
	GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan
}

type AdapterPegawaiService interface {
	ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionUpdateAbsenPulang(pegawai_id int, tgl string, pulang string, foto string)
	ActionInsertAbsenPulang(pegawai_id int, tgl string, pulang string, foto string)
	GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64)
	GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string)
	CheckHariLibur(unitkerja_id int, dayNow string) bool
	GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string)
	//-------------------------------------------------------------------------------------------------------
	CreatePerizinan(perizinan model.Perizinan) bool
	GetAllPerizinan(pegawai_id int) []model.APIResponsePerizinan
	GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan
	UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) bool
	//-------------------------------------------------------------------------------------------------------
	GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan
}
