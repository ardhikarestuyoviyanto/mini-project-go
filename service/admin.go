package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcAdmin struct {
	c    config.Config
	repo domain.AdapterAdminService
}

func (s *svcAdmin) UpdateRekapAbsen(starts string, finish string, pegawai_id int, rekapabsen model.RekapAbsen) {
	s.repo.UpdateRekapAbsen(starts, finish, pegawai_id, rekapabsen)
}

func (s *svcAdmin) DeleteRekapAbsen(starts string, finish string, pegawai_id int) {
	s.repo.DeleteRekapAbsen(starts, finish, pegawai_id)
}

func (s *svcAdmin) InsertRekapAbsen(rekapabsen model.RekapAbsen) {
	s.repo.InsertRekapAbsen(rekapabsen)
}

func (s *svcAdmin) DeletePerizinan(perizinan_id int) {
	s.repo.DeletePerizinan(perizinan_id)
}

func (s *svcAdmin) GetAllPerizinan() []model.APIResponsePerizinan {
	return s.repo.GetAllPerizinan()
}

func (s *svcAdmin) GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan {
	return s.repo.GetByIdPerizinan(perizinan_id)
}

func (s *svcAdmin) UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) {
	s.repo.UpdatePerizinan(perizinan_id, perizinan)
}

//--------------------------------------------------------------------------------------

func (s *svcAdmin) CreateKategoriPerizinan(perizinan model.KategoriPerizinan) {
	s.repo.CreateKategoriPerizinan(perizinan)
}

func (s *svcAdmin) DeleteKategoriPerizinan(kategori_id int) {
	s.repo.DeleteKategoriPerizinan(kategori_id)
}

func (s *svcAdmin) GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan {
	return s.repo.GetAllKategoriPerizinan()
}

func (s *svcAdmin) GetByIdKategoriPerizinan(kategori_id int) model.APIResponseKategoriPerizinan {
	return s.repo.GetByIdKategoriPerizinan(kategori_id)
}

func (s *svcAdmin) UpdateKategoriPerizinan(kategori_id int, perizinan model.KategoriPerizinan) {
	s.repo.UpdateKategoriPerizinan(kategori_id, perizinan)
}

//-----------------------------------------------------------------------------------------------

func (s *svcAdmin) GetByEmailPegawai(email string) model.APIResponsePegawai {
	return s.repo.GetByEmailPegawai(email)
}

func (s *svcAdmin) CreatePegawai(user model.Users) {
	s.repo.CreatePegawai(user)
}

func (s *svcAdmin) DeletePegawai(user_id int) {
	s.repo.DeletePegawai(user_id)
}

func (s *svcAdmin) GetAllPegawai() []model.APIResponsePegawai {
	return s.repo.GetAllPegawai()
}

func (s *svcAdmin) GetByIdPegawai(user_id int) model.APIResponsePegawai {
	return s.repo.GetByIdPegawai(user_id)
}

func (s *svcAdmin) UpdatePegawai(user_id int, user model.Users) {
	s.repo.UpdatePegawai(user_id, user)
}

//------------------------------------------------------------------------------------

func (s *svcAdmin) CreateUnitKerja(unitkerja model.UnitKerja) {
	s.repo.CreateUnitKerja(unitkerja)
}

func (s *svcAdmin) DeleteUnitKerja(unitkerja_id int) {
	s.repo.DeleteUnitKerja(unitkerja_id)
}

func (s *svcAdmin) GetAllUnitKerja() []model.APIResponseUnitKerja {
	return s.repo.GetAllUnitKerja()
}

func (s *svcAdmin) GetByIdUnitKerja(unitkerja_id int) model.APIResponseUnitKerja {
	return s.repo.GetByIdUnitKerja(unitkerja_id)
}

func (s *svcAdmin) UpdateUnitKerja(unitkerja_id int, unitkerja model.UnitKerja) {
	s.repo.UpdateUnitKerja(unitkerja_id, unitkerja)
}

//---------------------------------------------------------------------------------------------------------------

func (s *svcAdmin) GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail) {
	return s.repo.GetByIdJamKerjaDetail(jamkerja_id)
}

func (s *svcAdmin) CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{}) {
	s.repo.CreateJamKerjaDetail(jamkerjaDetail)
}

func (s *svcAdmin) DeleteJamKerjaDetail(jamkerja_id int) {
	s.repo.DeleteJamKerjaDetail(jamkerja_id)
}

//---------------------------------------------------------------------------------------------------------

func (s *svcAdmin) GetAllJamKerja() []model.APIResponseJamKerja {
	return s.repo.GetAllJamKerja()
}

func (s *svcAdmin) GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja {
	return s.repo.GetByIdJamKerja(jamkerja_id)
}

func (s *svcAdmin) DeleteJamKerja(jamkerja_id int) {
	s.repo.DeleteJamKerja(jamkerja_id)
}

func (s *svcAdmin) UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja) {
	s.repo.UpdateJamKerja(jamkerja_id, jamkerja)
}

func (s *svcAdmin) CreateJamKerja(jamkerja model.JamKerja) {
	s.repo.CreateJamKerja(jamkerja)
}

func NewServiceAdmin(repo domain.AdapterAdminRepository, c config.Config) domain.AdapterAdminService {
	return &svcAdmin{
		c:    c,
		repo: repo,
	}
}
