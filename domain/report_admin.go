package domain

import "mini-project-go/model"

type AdapterAdminReportRepository interface {
	GetBulanan(user_id, month int) []model.APIResponseRekapAbsen
	GetRangeDay(user_id int, start, finish string) []model.APIResponseRekapAbsen
}

type AdapterAdminReportServive interface {
	GetBulanan(user_id, month int) []model.APIResponseRekapAbsen
	GetRangeDay(user_id int, start, finish string) []model.APIResponseRekapAbsen
}
