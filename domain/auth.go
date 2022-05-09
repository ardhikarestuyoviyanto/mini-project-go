package domain

import "mini-project-go/model"

type AdapterAuthRepository interface {
	LoginUsers(email string, password string) (model.Users, bool)
	UpdateToken(token string, user_id int)
	LogoutUsers(user_id int) bool
	GetUsersByToken(token string) model.Users
}

type AdapterAuthService interface {
	LoginUsers(email string, password string) (model.Users, bool)
	UpdateToken(token string, user_id int)
	LogoutUsers(user_id int) bool
	GetUsersByToken(token string) model.Users
}
