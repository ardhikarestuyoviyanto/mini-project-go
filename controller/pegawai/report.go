package pegawai

import (
	"mini-project-go/constants"
	"mini-project-go/domain"
	"mini-project-go/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportEchoController struct {
	SvcAdmin domain.AdapterPegawaiReportServive
}

func (svc *ReportEchoController) GetBulanan(c echo.Context) error {

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]
	claimsToken, _ := middleware.DecodeTokenPegawai(token)

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	month, errMonth := dataRaw["bulan"].(float64)

	if !errMonth {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}
	report := svc.SvcAdmin.GetBulanan(int(claimsToken["user_id"].(float64)), int(month))
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    report,
		"total":   len(report),
	})
}

func (svc *ReportEchoController) GetRangeDay(c echo.Context) error {

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]
	claimsToken, _ := middleware.DecodeTokenPegawai(token)

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	start, errStart := dataRaw["start"].(string)
	finish, errFinish := dataRaw["finish"].(string)

	if !errStart || !errFinish {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}

	report := svc.SvcAdmin.GetRangeDay(int(claimsToken["user_id"].(float64)), start, finish)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    report,
		"success": true,
	})
}
