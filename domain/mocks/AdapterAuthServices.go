package mocks

import (
	"mini-project-go/model"

	"github.com/stretchr/testify/mock"
)

type SvcAuthMock struct {
	mock.Mock
}

func (s *SvcAuthMock) GetUsersByToken(token string) model.Users {
	return model.Users{
		Nama: "Ardhikaaaa",
	}
}

func (s *SvcAuthMock) UpdateToken(token string, user_id int) {
	s.Called(token, user_id)
}

func (s *SvcAuthMock) LogoutUsers(user_id int) bool {
	ret := s.Called(user_id)
	return ret.Get(0).(bool)
}

func (s *SvcAuthMock) LoginUsers(email string, password string) (model.Users, bool) {
	ret := s.Called(email, password)
	return ret.Get(0).(model.Users), ret.Get(1).(bool)
}
