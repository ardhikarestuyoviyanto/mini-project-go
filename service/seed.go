package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcSeed struct {
	c    config.Config
	repo domain.AdapterSeedService
}

func (s *svcSeed) MakePassword(password string) string {
	return s.repo.MakePassword(password)
}

func (s *svcSeed) CreateRole(role []model.Role) error {
	return s.repo.CreateRole(role)
}

func (s *svcSeed) CreateUserAdmin(users model.Users) error {
	return s.repo.CreateUserAdmin(users)
}

func NewServiceSeed(repo domain.AdapterSeedRepository, c config.Config) domain.AdapterSeedService {
	return &svcSeed{
		c:    c,
		repo: repo,
	}
}
