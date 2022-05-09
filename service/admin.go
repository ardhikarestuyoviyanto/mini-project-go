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