package mocks

import (
	"mini-project-go/model"

	"github.com/stretchr/testify/mock"
)

type SvcAdminMock struct {
	mock.Mock
}

func (s *SvcAdminMock) UpdateRekapAbsen(starts string, finish string, pegawai_id int, rekapabsen model.RekapAbsen) {
	s.Called(starts, finish, pegawai_id, rekapabsen)
}

func (s *SvcAdminMock) DeleteRekapAbsen(starts string, finish string, pegawai_id int) {
	s.Called(starts, finish, pegawai_id)
}

func (s *SvcAdminMock) InsertRekapAbsen(rekapabsen model.RekapAbsen) {
	s.Called(rekapabsen)
}

func (s *SvcAdminMock) DeletePerizinan(perizinan_id int) {
	s.Called(perizinan_id)
}

func (s *SvcAdminMock) GetAllPerizinan() []model.APIResponsePerizinan {
	return []model.APIResponsePerizinan{
		{
			ID:          1,
			PegawaiNama: "Ardhika Restu",
		},
	}
}

func (s *SvcAdminMock) GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan {
	return model.APIResponsePerizinan{
		KategoriPerizinanNama: "xxxx",
		PegawaiNama:           "Ardhika Restu",
	}
}

func (s *SvcAdminMock) UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) {
	s.Called(perizinan_id, perizinan)
}

//--------------------------------------------------------------------------------------

func (s *SvcAdminMock) CreateKategoriPerizinan(perizinan model.KategoriPerizinan) {
	s.Called(perizinan)
}

func (s *SvcAdminMock) DeleteKategoriPerizinan(kategori_id int) {
	s.Called(kategori_id)
}

func (s *SvcAdminMock) GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan {
	return []model.APIResponseKategoriPerizinan{
		{
			ID:   1,
			Name: "Sakit",
		},
	}
}

func (s *SvcAdminMock) GetByIdKategoriPerizinan(kategori_id int) model.APIResponseKategoriPerizinan {
	return model.APIResponseKategoriPerizinan{
		Name: "Sakit",
		ID:   1,
	}
}

func (s *SvcAdminMock) UpdateKategoriPerizinan(kategori_id int, perizinan model.KategoriPerizinan) {
	s.Called(kategori_id, perizinan)
}

//-----------------------------------------------------------------------------------------------

func (s *SvcAdminMock) GetByEmailPegawai(email string) model.APIResponsePegawai {
	return model.APIResponsePegawai{
		Nama:  "Ardhika Restu Yoviyanto",
		Email: "asasas@gmail.com",
	}
}

func (s *SvcAdminMock) CreatePegawai(user model.Users) {
	s.Called(user)
}

func (s *SvcAdminMock) DeletePegawai(user_id int) {
	s.Called(user_id)
}

func (s *SvcAdminMock) GetAllPegawai() []model.APIResponsePegawai {
	return []model.APIResponsePegawai{
		{
			Nama: "Ardhika Restu Yoviyanto",
		},
	}
}

func (s *SvcAdminMock) GetByIdPegawai(user_id int) model.APIResponsePegawai {
	return model.APIResponsePegawai{
		Nama: "Ardhika Restu Yoviyanto",
	}
}

func (s *SvcAdminMock) UpdatePegawai(user_id int, user model.Users) {
	s.Called(user_id, user)
}

//------------------------------------------------------------------------------------

func (s *SvcAdminMock) CreateUnitKerja(unitkerja model.UnitKerja) {
	s.Called(unitkerja)
}

func (s *SvcAdminMock) DeleteUnitKerja(unitkerja_id int) {
	s.Called(unitkerja_id)
}

func (s *SvcAdminMock) GetAllUnitKerja() []model.APIResponseUnitKerja {
	return []model.APIResponseUnitKerja{
		{
			Nama: "Unit Kerja X",
		},
	}
}

func (s *SvcAdminMock) GetByIdUnitKerja(unitkerja_id int) model.APIResponseUnitKerja {
	return model.APIResponseUnitKerja{
		Nama: "Unit Kerja X",
	}
}

func (s *SvcAdminMock) UpdateUnitKerja(unitkerja_id int, unitkerja model.UnitKerja) {
	s.Called(unitkerja_id, unitkerja)
}

//---------------------------------------------------------------------------------------------------------------

func (s *SvcAdminMock) GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail) {
	return model.APIResponseJamKerja{}, []model.APIResponseJamKerjaDetail{}
}

func (s *SvcAdminMock) CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{}) {
	s.Called(jamkerjaDetail)
}

func (s *SvcAdminMock) DeleteJamKerjaDetail(jamkerja_id int) {
	s.Called(jamkerja_id)
}

//---------------------------------------------------------------------------------------------------------

func (s *SvcAdminMock) GetAllJamKerja() []model.APIResponseJamKerja {
	return []model.APIResponseJamKerja{}
}

func (s *SvcAdminMock) GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja {
	return model.APIResponseJamKerja{}
}

func (s *SvcAdminMock) DeleteJamKerja(jamkerja_id int) {
	s.Called(jamkerja_id)
}

func (s *SvcAdminMock) UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja) {
	s.Called(jamkerja_id, jamkerja)
}

func (s *SvcAdminMock) CreateJamKerja(jamkerja model.JamKerja) {
	s.Called(jamkerja)
}
