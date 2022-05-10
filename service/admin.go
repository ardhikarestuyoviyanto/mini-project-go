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
