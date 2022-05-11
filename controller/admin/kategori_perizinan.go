package admin

import (
	"mini-project-go/domain"
	"mini-project-go/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KategoriPerizinanEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *KategoriPerizinanEchoController) CreateController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	name, errName := dataRaw["name"].(string)
	maxDay, errMaxDay := dataRaw["max_day"].(float64)
	types, errTypes := dataRaw["type"].(string)

	if !errName || !errMaxDay || !errTypes {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
			"success": false,
		})
	}

	if types != "izin" && types != "cuti" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "type perizinan must be enum : izin or cuti",
			"success": false,
		})
	}

	perizinan := model.KategoriPerizinan{
		Name:   name,
		MaxDay: int(maxDay),
		Type:   types,
	}

	svc.SvcAdmin.CreateKategoriPerizinan(perizinan)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "kategori perizinan berhasil ditambahkan",
		"success": true,
	})
}

func (svc *KategoriPerizinanEchoController) GetAllController(c echo.Context) error {
	perizinan := svc.SvcAdmin.GetAllKategoriPerizinan()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    perizinan,
		"total":   len(perizinan),
		"success": true,
	})
}

func (svc *KategoriPerizinanEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters id",
			"success": false,
		})
	}
	perizinan := svc.SvcAdmin.GetByIdKategoriPerizinan(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    perizinan,
		"success": true,
	})
}

func (svc *KategoriPerizinanEchoController) UpdateController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters id",
			"success": false,
		})
	}
	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	name, errName := dataRaw["name"].(string)
	maxDay, errMaxDay := dataRaw["max_day"].(float64)
	types, errTypes := dataRaw["type"].(string)

	if !errName || !errMaxDay || !errTypes {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
			"success": false,
		})
	}

	if types != "izin" && types != "cuti" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "type perizinan must be enum : izin or cuti",
			"success": false,
		})
	}

	perizinan := model.KategoriPerizinan{
		Name:   name,
		MaxDay: int(maxDay),
		Type:   types,
	}
	svc.SvcAdmin.UpdateKategoriPerizinan(id, perizinan)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "kategori perizinan berhasil diupdate",
	})
}

func (svc *KategoriPerizinanEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters id",
			"success": false,
		})
	}
	svc.SvcAdmin.DeleteKategoriPerizinan(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "kategori perizinan berhasil dihapus",
	})
}
