package domain

import "mini-project-go/model"

type AdapterSeedRepository interface {
	CreateUserAdmin(users model.Users) error
	MakePassword(password string) string
	CreateRole(role []model.Role) error
}

type AdapterSeedService interface {
	CreateUserAdmin(users model.Users) error
	MakePassword(password string) string
	CreateRole(role []model.Role) error
}
