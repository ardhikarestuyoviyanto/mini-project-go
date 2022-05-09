package repository

import (
	"mini-project-go/domain"
	"mini-project-go/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repositoryAuth struct {
	DB *gorm.DB
}

func (r *repositoryAuth) GetUsersByToken(token string) model.Users {
	users := model.Users{}
	r.DB.First(&users, "token = ?", token)
	return users
}

func (r *repositoryAuth) UpdateToken(token string, user_id int) {
	users := model.Users{}
	r.DB.Model(&users).Where("id", user_id).Update("token", token)
}

func (r *repositoryAuth) LogoutUsers(user_id int) bool {
	res := r.DB.Model(&model.Users{}).Where("id", user_id).Update("token", "")
	if res != nil {
		return false
	} else {
		return true
	}
}

func (r *repositoryAuth) LoginUsers(email string, password string) (model.Users, bool) {
	users := model.Users{}
	err_email := r.DB.Where("email = ?", email).First(&users).Error
	if err_email != nil {
		return users, false
	}
	err_password := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	if err_password != nil && err_password == bcrypt.ErrMismatchedHashAndPassword {
		return users, false
	}
	return users, true
}

func NewAuthRepository(db *gorm.DB) domain.AdapterAuthRepository {
	return &repositoryAuth{
		DB: db,
	}
}
