package database

import (
	"mini-project-go/config"
	"mini-project-go/model"
	"mini-project-go/repository"
	"mini-project-go/service"

	"gorm.io/gorm"
)

func CreateRoleSeed(db *gorm.DB, conf config.Config) {
	repo := repository.NewSeedRepository(db)
	serv := service.NewServiceSeed(repo, conf)

	role := []model.Role{
		{
			ID:       1,
			RoleName: "admin",
		},
		{
			ID:       2,
			RoleName: "pegawai",
		},
	}

	serv.CreateRole(role)
}

func CreateAdminSeed(db *gorm.DB, conf config.Config) {
	repo := repository.NewSeedRepository(db)
	serv := service.NewServiceSeed(repo, conf)

	user := model.Users{
		Nama:     "Ardhika Yoviyanto, S.kom",
		Email:    "ardhikayoviyanto@gmail.com",
		Password: serv.MakePassword("123"),
		Jabatan:  "Administrator",
		RoleID:   1,
	}

	serv.CreateUserAdmin(user)
}
