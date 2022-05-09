package admin

import (
	"mini-project-go/domain"
	"mini-project-go/lib"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JamKerjaDetailEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *JamKerjaDetailEchoController) CreateController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	jamkerja_id, errJamKerjaId := dataRaw["jamkerja_id"].(float64)
	data, errData := dataRaw["data"].([]interface{})

	if !errJamKerjaId || !errData || len(data) != 7 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form input",
		})
	}

	jamkerjaDetail := lib.CreateSliceJamKerjaDetail(data, int(jamkerja_id))

	if !lib.ValidateDayJamKerjaDetail(jamkerjaDetail) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid day name in form , ex : [senin, selasa, rabu, kamis, jumat, sabtu, minggu]",
		})
	}

	svc.SvcAdmin.DeleteJamKerjaDetail(int(jamkerja_id))
	svc.SvcAdmin.CreateJamKerjaDetail(jamkerjaDetail)

	jamkerja, jamkerjaDetails := svc.SvcAdmin.GetByIdJamKerjaDetail(int(jamkerja_id))

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Detail jam kerja berhasil diupdate",
		"data": map[string]interface{}{
			"jamkerja":        jamkerja,
			"jamkerja_detail": jamkerjaDetails,
		},
	})

}

func (svc *JamKerjaDetailEchoController) GetByIdController(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}

	jamkerja, jamkerjaDetails := svc.SvcAdmin.GetByIdJamKerjaDetail(idInt)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"jamkerja":        jamkerja,
			"jamkerja_detail": jamkerjaDetails,
		},
	})
}
