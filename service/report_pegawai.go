package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcReportPegawai struct {
	c    config.Config
	repo domain.AdapterPegawaiReportServive
}

func (s *svcReportPegawai) GetBulanan(user_id int, month int) []model.APIResponseRekapAbsen {
	return s.repo.GetBulanan(user_id, month)
}

func (s *svcReportPegawai) GetRangeDay(user_id int, start string, finish string) []model.APIResponseRekapAbsen {
	return s.repo.GetRangeDay(user_id, start, finish)
}

func NewServiceReportPegawai(repo domain.AdapterPegawaiReportRepository, c config.Config) domain.AdapterPegawaiReportServive {
	return &svcReportPegawai{
		c:    c,
		repo: repo,
	}
}
