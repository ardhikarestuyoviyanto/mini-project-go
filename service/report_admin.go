package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcReportAdmin struct {
	c    config.Config
	repo domain.AdapterAdminReportServive
}

func (s *svcReportAdmin) GetRangeDay(user_id int, start string, finish string) []model.APIResponseRekapAbsen {
	return s.repo.GetRangeDay(user_id, start, finish)
}

func (s *svcReportAdmin) GetBulanan(user_id int, month int) []model.APIResponseRekapAbsen {
	return s.repo.GetBulanan(user_id, month)
}

func NewServiceReportAdmin(repo domain.AdapterAdminReportRepository, c config.Config) domain.AdapterAdminReportServive {
	return &svcReportAdmin{
		c:    c,
		repo: repo,
	}
}
