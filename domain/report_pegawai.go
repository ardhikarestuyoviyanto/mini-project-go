package domain

import "mini-project-go/model"

type AdapterPegawaiReportRepository interface {
	GetBulanan(user_id, month int) []model.APIResponseRekapAbsen
	GetRangeDay(user_id int, start, finish string) []model.APIResponseRekapAbsen
}

type AdapterPegawaiReportServive interface {
	GetBulanan(user_id, month int) []model.APIResponseRekapAbsen
	GetRangeDay(user_id int, start, finish string) []model.APIResponseRekapAbsen
}
