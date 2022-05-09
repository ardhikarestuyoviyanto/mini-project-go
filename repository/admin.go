package repository

import (
	"mini-project-go/domain"
	"mini-project-go/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repositoryAdmin struct {
	DB *gorm.DB
}

func (r *repositoryAdmin) GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail) {
	var jamkerjaDetail []model.APIResponseJamKerjaDetail
	var jamkerja model.APIResponseJamKerja
	r.DB.Table("jamkerja").Where("id", jamkerja_id).Scan(&jamkerja)
	r.DB.Table("jamkerja_detail").Where("jamkerja_id", jamkerja_id).Scan(&jamkerjaDetail)
	return jamkerja, jamkerjaDetail
}

func (r *repositoryAdmin) CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{}) {
	r.DB.Model(&model.JamKerjaDetail{}).Create(jamkerjaDetail)
}

func (r *repositoryAdmin) DeleteJamKerjaDetail(jamkerja_id int) {
	r.DB.Unscoped().Where("jamkerja_id", jamkerja_id).Delete(&model.JamKerjaDetail{})
}

//----------------------------------------------------------------------------

func (r *repositoryAdmin) DeleteJamKerja(jamkerja_id int) {
	r.DB.Unscoped().Delete(&model.JamKerja{}, jamkerja_id)
}

func (r *repositoryAdmin) GetAllJamKerja() []model.APIResponseJamKerja {
	var jamkerja []model.APIResponseJamKerja
	r.DB.Table("jamkerja").Select("id", "nama", "harilibur").Scan(&jamkerja)
	return jamkerja
}

func (r *repositoryAdmin) GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja {
	var jamkerja model.APIResponseJamKerja
	r.DB.Table("jamkerja").Select("id", "nama", "harilibur").Where("id", jamkerja_id).Scan(&jamkerja)
	return jamkerja
}

func (r *repositoryAdmin) UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja) {
	r.DB.Model(&model.JamKerja{}).Select("nama", "harilibur").Where("id", jamkerja_id).Updates(jamkerja)
}

func (r *repositoryAdmin) CreateJamKerja(jamkerja model.JamKerja) {
	r.DB.Create(&jamkerja)
}

//----------------------------------------------------------------------

func (r *repositoryAdmin) CreateRole(role []model.Role) error {
	r.DB.Create(&role)
	return nil
}

func (r *repositoryAdmin) MakePassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (r *repositoryAdmin) CreateUserAdmin(users model.Users) error {
	r.DB.Create(&users)
	return nil
}

func NewSeedRepository(db *gorm.DB) domain.AdapterSeedRepository {
	return &repositoryAdmin{
		DB: db,
	}
}

func NewAdminRepository(db *gorm.DB) domain.AdapterAdminRepository {
	return &repositoryAdmin{
		DB: db,
	}
}
