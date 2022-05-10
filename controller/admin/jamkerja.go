package admin

import (
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type JamKerjaEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *JamKerjaEchoController) CreateController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	namaStr, errNama := dataRaw["nama"].(string)
	hariLibur, errHariLibur := dataRaw["harilibur"].([]interface{})

	if !errNama || !errHariLibur || !lib.CheckDataType(hariLibur) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form",
		})
	}

	if !lib.ValidateDay(lib.InterfaceToSliceStr(hariLibur)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "nama harilibur invalid, ex : [senin, selasa, rabu, kamis, jumat, sabtu, minggu]",
		})
	}

	hariLiburStr := strings.Join(lib.InterfaceToSliceStr(hariLibur), ",")

	jamkerjaModel := model.JamKerja{
		Nama:      namaStr,
		Harilibur: hariLiburStr,
	}

	svc.SvcAdmin.CreateJamKerja(jamkerjaModel)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Jam kerja baru berhasil disimpan",
	})
}

func (svc *JamKerjaEchoController) GetAllController(c echo.Context) error {
	res := svc.SvcAdmin.GetAllJamKerja()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"total":   len(res),
		"data":    res,
	})
}

func (svc *JamKerjaEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}
	res := svc.SvcAdmin.GetByIdJamKerja(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    res,
	})
}

func (svc *JamKerjaEchoController) UpdateController(c echo.Context) error {

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

	namaStr, errNama := dataRaw["nama"].(string)
	hariLibur, errHariLibur := dataRaw["harilibur"].([]interface{})

	if !errNama || !errHariLibur || !lib.CheckDataType(hariLibur) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form",
		})
	}

	if !lib.ValidateDay(lib.InterfaceToSliceStr(hariLibur)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "nama harilibur invalid, ex : [senin, selasa, rabu, kamis, jumat, sabtu, minggu]",
		})
	}

	hariLiburStr := strings.Join(lib.InterfaceToSliceStr(hariLibur), ",")

	jamkerjaModel := model.JamKerja{
		Nama:      namaStr,
		Harilibur: hariLiburStr,
	}

	svc.SvcAdmin.UpdateJamKerja(id, jamkerjaModel)
	res := svc.SvcAdmin.GetByIdJamKerja(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "jam kerja berhasil diupdate",
		"data":    res,
	})
}

func (svc *JamKerjaEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameter id",
		})
	}

	svc.SvcAdmin.DeleteJamKerja(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "jam kerja berhasil dihapus",
	})
}

//--------------------------------------------------------------------------
