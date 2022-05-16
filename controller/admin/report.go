package admin

import (
	"mini-project-go/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportEchoController struct {
	SvcAdmin domain.AdapterAdminReportServive
}

func (svc *ReportEchoController) GetBulanan(c echo.Context) error {
	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	userId, errUser := dataRaw["user_id"].(float64)
	month, errMonth := dataRaw["bulan"].(float64)
	if !errUser || !errMonth {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}
	report := svc.SvcAdmin.GetBulanan(int(userId), int(month))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    report,
		"total":   len(report),
	})
}

func (svc *ReportEchoController) GetRangeDay(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	userId, errUser := dataRaw["user_id"].(float64)
	start, errStart := dataRaw["start"].(string)
	finish, errFinish := dataRaw["finish"].(string)

	if !errUser || !errStart || !errFinish {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}

	report := svc.SvcAdmin.GetRangeDay(int(userId), start, finish)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    report,
		"success": true,
	})
}
