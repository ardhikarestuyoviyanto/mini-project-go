package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcAuth struct {
	c    config.Config
	repo domain.AdapterAuthService
}

func (s *svcAuth) GetUsersByToken(token string) model.Users {
	return s.repo.GetUsersByToken(token)
}

func (s *svcAuth) UpdateToken(token string, user_id int) {
	s.repo.UpdateToken(token, user_id)
}

func (s *svcAuth) LogoutUsers(user_id int) bool {
	return s.repo.LogoutUsers(user_id)
}

func (s *svcAuth) LoginUsers(email string, password string) (model.Users, bool) {
	return s.repo.LoginUsers(email, password)
}

func NewServiceAuth(repo domain.AdapterAuthRepository, c config.Config) domain.AdapterAuthService {
	return &svcAuth{
		c:    c,
		repo: repo,
	}
}
