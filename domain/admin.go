package domain

import (
	"mini-project-go/model"
)

type AdapterAdminRepository interface {
	CreateJamKerja(jamkerja model.JamKerja)
	UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja)
	GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja
	GetAllJamKerja() []model.APIResponseJamKerja
	DeleteJamKerja(jamkerja_id int)
	//------------------------------------------------------
	CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{})
	DeleteJamKerjaDetail(jamkerja_id int)
	GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail)
}

type AdapterAdminService interface {
	CreateJamKerja(jamkerja model.JamKerja)
	UpdateJamKerja(jamkerja_id int, jamkerja model.JamKerja)
	GetByIdJamKerja(jamkerja_id int) model.APIResponseJamKerja
	GetAllJamKerja() []model.APIResponseJamKerja
	DeleteJamKerja(jamkerja_id int)
	//---------------------------------------------------------
	CreateJamKerjaDetail(jamkerjaDetail []map[string]interface{})
	DeleteJamKerjaDetail(jamkerja_id int)
	GetByIdJamKerjaDetail(jamkerja_id int) (model.APIResponseJamKerja, []model.APIResponseJamKerjaDetail)
}
