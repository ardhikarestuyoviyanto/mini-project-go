package domain

type AdapterPegawaiRepository interface {
	ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64)
	GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string)
	CheckHariLibur(unitkerja_id int, dayNow string) bool
	GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string)
}

type AdapterPegawaiService interface {
	ActionUpdateAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	ActionInsertAbsenMasuk(pegawai_id int, tgl string, masuk string, foto string)
	GetLatiduteLongtiduteUnitKerja(unitkerja_id int) (latiduteUK float64, longtidteUK float64)
	GetJamKerjaDetailTodayByIdUnitKerja(unitkerja_id int, hari string) (mMasuk string, bMasuk string, mPulang string, bPulang string)
	CheckHariLibur(unitkerja_id int, dayNow string) bool
	GetPresensiToday(user_id int, tanggal string) (absenMasuk string, absenPulang string)
}
