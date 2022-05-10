package domain

import "mini-project-go/model"

type AdapterSeedRepository interface {
	CreateUserAdmin(users model.Users) error
	CreateRole(role []model.Role) error
}

type AdapterSeedService interface {
	CreateUserAdmin(users model.Users) error
	CreateRole(role []model.Role) error
}
