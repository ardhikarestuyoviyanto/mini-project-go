package repository

import (
	"mini-project-go/domain"
	"mini-project-go/model"

	"gorm.io/gorm"
)

type repositoryAdmin struct {
	DB *gorm.DB
}

func (r *repositoryAdmin) GetByEmailPegawai(email string) model.APIResponsePegawai {
	var pegawai model.APIResponsePegawai
	r.DB.Table("users").Select("users.id", "users.nama", "email", "jabatan", "unitkerja_id", "unitkerja.nama AS unitkerja_nama").Joins("inner join unitkerja on unitkerja.id = users.unitkerja_id").Where(map[string]interface{}{"email": email, "role_id": 2}).Scan(&pegawai)
	return pegawai
}

func (r *repositoryAdmin) CreatePegawai(user model.Users) {
	r.DB.Create(&user)
}

func (r *repositoryAdmin) DeletePegawai(user_id int) {
	r.DB.Unscoped().Where("id", user_id).Delete(&model.Users{})
}

func (r *repositoryAdmin) GetAllPegawai() []model.APIResponsePegawai {
	var pegawai []model.APIResponsePegawai
	r.DB.Table("users").Select("users.id", "users.nama", "email", "jabatan", "unitkerja_id", "unitkerja.nama AS unitkerja_nama").Joins("inner join unitkerja on unitkerja.id = users.unitkerja_id").Where("role_id", 2).Scan(&pegawai)
	return pegawai
}

func (r *repositoryAdmin) GetByIdPegawai(user_id int) model.APIResponsePegawai {
	var pegawai model.APIResponsePegawai
	r.DB.Table("users").Select("users.id", "users.nama", "email", "jabatan", "unitkerja_id", "unitkerja.nama AS unitkerja_nama").Joins("inner join unitkerja on unitkerja.id = users.unitkerja_id").Where(map[string]interface{}{"users.id": user_id, "role_id": 2}).Scan(&pegawai)
	return pegawai
}

func (r *repositoryAdmin) UpdatePegawai(user_id int, user model.Users) {
	r.DB.Model(&model.Users{}).Where("id", user_id).Updates(user)
}

//----------------------------------------------------------------------------------------------------------------------

func (r *repositoryAdmin) CreateUnitKerja(unitkerja model.UnitKerja) {
	r.DB.Create(&unitkerja)
}

func (r *repositoryAdmin) DeleteUnitKerja(unitkerja_id int) {
	r.DB.Unscoped().Where("id", unitkerja_id).Delete(&model.UnitKerja{})
}

func (r *repositoryAdmin) GetAllUnitKerja() []model.APIResponseUnitKerja {
	var unitkerja []model.APIResponseUnitKerja
	r.DB.Table("unitkerja").Select("unitkerja.id", "unitkerja.nama", "alamat", "latidute", "longtidute", "jamkerja_id", "jamkerja.nama AS jamkerja_nama").Joins("inner join jamkerja on jamkerja.id = unitkerja.jamkerja_id ").Scan(&unitkerja)
	return unitkerja
}

func (r *repositoryAdmin) GetByIdUnitKerja(unitkerja_id int) model.APIResponseUnitKerja {
	var unitkerja model.APIResponseUnitKerja
	r.DB.Table("unitkerja").Select("unitkerja.id", "unitkerja.nama", "alamat", "latidute", "longtidute", "jamkerja_id", "jamkerja.nama AS jamkerja_nama").Joins("inner join jamkerja on jamkerja.id = unitkerja.jamkerja_id ").Where("unitkerja.id", unitkerja_id).Scan(&unitkerja)
	return unitkerja
}

func (r *repositoryAdmin) UpdateUnitKerja(unitkerja_id int, unitkerja model.UnitKerja) {
	r.DB.Model(&model.UnitKerja{}).Where("id", unitkerja_id).Updates(unitkerja)
}

//-------------------------------------------------------------------------------------------------------------------------------

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
