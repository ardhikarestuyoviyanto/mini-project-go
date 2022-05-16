package service

import (
	"mini-project-go/config"
	"mini-project-go/domain"
	"mini-project-go/model"
)

type svcPegawai struct {
	c    config.Config
	repo domain.AdapterPegawaiRepository
}

func (s *svcPegawai) CheckPerizinanDay(user_id int, tgl string) bool {
	return s.repo.CheckPerizinanDay(user_id, tgl)
}

func (s *svcPegawai) GetAllKategoriPerizinan() []model.APIResponseKategoriPerizinan {
	return s.repo.GetAllKategoriPerizinan()
}

//--------------------------------------------------------------------------------------------
func (s *svcPegawai) CreatePerizinan(perizinan model.Perizinan) bool {
	return s.repo.CreatePerizinan(perizinan)
}

func (s *svcPegawai) GetAllPerizinan(pegawai_id int) []model.APIResponsePerizinan {
	return s.repo.GetAllPerizinan(pegawai_id)
}

func (s *svcPegawai) GetByIdPerizinan(perizinan_id int) model.APIResponsePerizinan {
	return s.repo.GetByIdPerizinan(perizinan_id)
}

func (s *svcPegawai) UpdatePerizinan(perizinan_id int, perizinan model.Perizinan) bool {
	return s.repo.UpdatePerizinan(perizinan_id, perizinan)
}

//-----------------------------------------------------------------------------------------------------

func (s *svcPegawai) ActionInsertAbsenPulang(pegawai_id int, tgl string, pulang string, foto string) {
	s.repo.ActionInsertAbsenPulang(pegawai_id, tgl, pulang, foto)
}

func (s *svcPegawai) ActionUpdateAbsenPulang(pegawai_id int, tgl string, pulang string, foto string) {
	s.repo.ActionUpdateAbsenPulang(pegawai_id, tgl, pulang, foto)
}

func (s *svcPegawai) ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string) {
	s.repo.ActionInsertAbsenMasuk(pegawai_id, tgl, masuk, foto)
}

func (s *svcPegawai) GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string) {
	return s.repo.GetPresensiToday(user_id, tanggal)
}

func (s *svcPegawai) CheckHariLibur(unitkerja_id int, dayNow string) bool {
	return s.repo.CheckHariLibur(unitkerja_id, dayNow)
}

func (s *svcPegawai) GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string) {
	return s.repo.GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id, hari)
}

func (s *svcPegawai) GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64) {
	return s.repo.GetLatiduteLongtiduteUnitKerja(unitkerja_id)
}

func (s *svcPegawai) ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string) {
	s.repo.ActionUpdateAbsenMasuk(pegawai_id, tgl, masuk, foto)
}

func NewServicePegawai(repo domain.AdapterPegawaiRepository, c config.Config) domain.AdapterPegawaiService {
	return &svcPegawai{
		c:    c,
		repo: repo,
	}
}
