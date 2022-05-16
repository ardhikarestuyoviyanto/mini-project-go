package domain

import (
	"mini-project-go/model"
)

type AdapterAdminRepository interface {
	CreateJamKerja(jamkerja model.JamKerja)
	UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja)
	GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja
	GetAllJamKerja() []model.APIResponseJamKerja
	DeleteJamKerja(jamkerja_id int)
	//------------------------------------------------------
	CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{})
	DeleteJamKerjaDetail(jamkerja_id int)
	GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail)
	//-----------------------------------------------------------
	CreateUnitKerja(unitkerja model.UnitKerja)
	UpdateUnitKerja(unitkerja_id int, unitkerja model.UnitKerja)
	GetAllUnitKerja() []model.APIResponseUnitKerja
	GetByIdUnitKerja(unitkerja_id int) model.APIResponseUnitKerja
	DeleteUnitKerja(unitkerja_id int)
	//------------------------------------------------------------
	CreatePegawai(user model.Users)
	UpdatePegawai(user_id int, user model.Users)
	GetAllPegawai() []model.APIResponsePegawai
	GetByIdPegawai(user_id int) model.APIResponsePegawai
	DeletePegawai(user_id int)
	GetByEmailPegawai(email string) model.APIResponsePegawai
	//------------------------------------------------------------
	CreateKategoriPerizinan(perizinan model.KategoriPerizinan)
	UpdateKategoriPerizinan(kategori_id int, perizinan model.KategoriPerizinan)
	DeleteKategoriPerizinan(kategori_id int)
	GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan
	GetByIdKategoriPerizinan(kategori_id int) model.APIResponseKategoriPerizinan
	//-----------------------------------------------------------------------
	GetAllPerizinan() []model.APIResponsePerizinan
	GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan
	UpdatePerizinan(perizinan_id int, perizinan model.Perizinan)
	InsertRekapAbsen(rekapabsen model.RekapAbsen)
	DeleteRekapAbsen(starts, finish string, pegawai_id int)
	UpdateRekapAbsen(starts, finish string, pegawai_id int, rekapabsen model.RekapAbsen)
	DeletePerizinan(perizinan_id int)
}

type AdapterAdminService interface {
	CreateJamKerja(jamkerja model.JamKerja)
	UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja)
	GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja
	GetAllJamKerja() []model.APIResponseJamKerja
	DeleteJamKerja(jamkerja_id int)
	//---------------------------------------------------------
	CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{})
	DeleteJamKerjaDetail(jamkerja_id int)
	GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail)
	//-----------------------------------------------------------
	CreateUnitKerja(unitkerja model.UnitKerja)
	UpdateUnitKerja(unitkerja_id int, unitkerja model.UnitKerja)
	GetAllUnitKerja() []model.APIResponseUnitKerja
	GetByIdUnitKerja(unitkerja_id int) model.APIResponseUnitKerja
	DeleteUnitKerja(unitkerja_id int)
	//----------------------------------------------------------------
	CreatePegawai(user model.Users)
	UpdatePegawai(user_id int, user model.Users)
	GetAllPegawai() []model.APIResponsePegawai
	GetByIdPegawai(user_id int) model.APIResponsePegawai
	DeletePegawai(user_id int)
	GetByEmailPegawai(email string) model.APIResponsePegawai
	//----------------------------------------------------------------
	CreateKategoriPerizinan(perizinan model.KategoriPerizinan)
	UpdateKategoriPerizinan(kategori_id int, perizinan model.KategoriPerizinan)
	DeleteKategoriPerizinan(kategori_id int)
	GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan
	GetByIdKategoriPerizinan(kategori_id int) model.APIResponseKategoriPerizinan
	//--------------------------------------------------------------------------
	GetAllPerizinan() []model.APIResponsePerizinan
	GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan
	UpdatePerizinan(perizinan_id int, perizinan model.Perizinan)
	InsertRekapAbsen(rekapabsen model.RekapAbsen)
	DeleteRekapAbsen(starts, finish string, pegawai_id int)
	UpdateRekapAbsen(starts, finish string, pegawai_id int, rekapabsen model.RekapAbsen)
	DeletePerizinan(perizinan_id int)
}
