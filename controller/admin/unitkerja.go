package admin

import (
	"mini-project-go/domain"
	"mini-project-go/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UnitKerjaEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *UnitKerjaEchoController) CreateController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	nama, errName := dataRaw["nama"].(string)
	alamat, errAlamat := dataRaw["alamat"].(string)
	lat, errLat := dataRaw["latidute"].(float64)
	long, errLong := dataRaw["longtidute"].(float64)
	jamkerjaId, errJamkerjaId := dataRaw["jamkerja_id"].(float64)

	if !errName || !errAlamat || !errLat || !errLong || !errJamkerjaId || len(nama) == 0 || len(alamat) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form input",
		})
	}

	unitkerja := model.UnitKerja{
		Nama:       nama,
		Alamat:     alamat,
		Latidute:   lat,
		Longtidute: long,
		JamkerjaID: int(jamkerjaId),
	}

	svc.SvcAdmin.CreateUnitKerja(unitkerja)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Unit kerja baru berhasil disimpan",
	})
}

func (svc *UnitKerjaEchoController) GetAllController(c echo.Context) error {
	unitkerja := svc.SvcAdmin.GetAllUnitKerja()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"total":   len(unitkerja),
		"data":    unitkerja,
	})
}

func (svc *UnitKerjaEchoController) GetByIdController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}

	unitkerja := svc.SvcAdmin.GetByIdUnitKerja(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    unitkerja,
		"success": true,
	})
}

func (svc *UnitKerjaEchoController) UpdateController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	nama, errName := dataRaw["nama"].(string)
	alamat, errAlamat := dataRaw["alamat"].(string)
	lat, errLat := dataRaw["latidute"].(float64)
	long, errLong := dataRaw["longtidute"].(float64)
	jamkerjaId, errJamkerjaId := dataRaw["jamkerja_id"].(float64)

	if !errName || !errAlamat || !errLat || !errLong || !errJamkerjaId || len(nama) == 0 || len(alamat) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form input",
		})
	}

	unitkerja := model.UnitKerja{
		Nama:       nama,
		Alamat:     alamat,
		Latidute:   lat,
		Longtidute: long,
		JamkerjaID: int(jamkerjaId),
	}

	svc.SvcAdmin.UpdateUnitKerja(id, unitkerja)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Unit kerja berhasil diupdate",
	})

}

func (svc *UnitKerjaEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}
	svc.SvcAdmin.DeleteUnitKerja(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "Unit kerja berhasil dihapus",
	})

}
