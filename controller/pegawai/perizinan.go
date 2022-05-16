package pegawai

import (
	"io"
	"mini-project-go/constants"
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/middleware"
	"mini-project-go/model"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PerizinanEchoController struct {
	SvcPegawai domain.AdapterPegawaiService
}

func (svc *PerizinanEchoController) CreateController(c echo.Context) error {

	start := c.FormValue("start")
	finish := c.FormValue("finish")
	kategori_perizinan_id, errId := strconv.Atoi(c.FormValue("kategori_perizinan_id"))
	file_pendukung, errFile := c.FormFile("file_pendukung")

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]
	claimsToken, _ := middleware.DecodeTokenPegawai(token)

	if errFile != nil || errId != nil || len(start) == 0 || len(finish) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
			"success": false,
		})
	}

	pathFileInServer := os.Getenv("BASE_URL") + constants.STATIC_FILE_PERIZINAN + file_pendukung.Filename

	if !lib.CheckExtensionForPdf(filepath.Ext(pathFileInServer)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid extension file (extenion must be .pdf)",
		})
	}

	src, err := file_pendukung.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	dir, _ := os.Getwd()

	locationFile := filepath.Join(dir, constants.DIR_FILE_PERIZINAN, file_pendukung.Filename)
	dst, err := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	perizinan := model.Perizinan{
		Status:              "P",
		Start:               start,
		Finish:              finish,
		FilePendukung:       pathFileInServer,
		KategoriPerizinanID: kategori_perizinan_id,
		UserID:              int(claimsToken["user_id"].(float64)),
	}

	if !svc.SvcPegawai.CreatePerizinan(perizinan) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "maximum hari yang diperbolehkan mengajukan perizinan ini melebihi batas range izin",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Pengajuan perizinan anda berhasil disimpan, silahkan menunggu admin meng-acc perizinan anda",
	})
}

func (svc *PerizinanEchoController) UpdateController(c echo.Context) error {

	start := c.FormValue("start")
	finish := c.FormValue("finish")
	kategori_perizinan_id, errIdKategori := strconv.Atoi(c.FormValue("kategori_perizinan_id"))
	perizinan_id, errIdPerizinan := strconv.Atoi(c.Param("id"))
	file_pendukung, errFile := c.FormFile("file_pendukung")

	if errFile != nil || errIdKategori != nil || len(start) == 0 || len(finish) == 0 || errIdPerizinan != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form",
			"success": false,
		})
	}

	pathFileInServer := os.Getenv("BASE_URL") + constants.STATIC_FILE_PERIZINAN + file_pendukung.Filename

	if !lib.CheckExtensionForPdf(filepath.Ext(pathFileInServer)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid extension file (extenion must be .pdf)",
		})
	}

	src, err := file_pendukung.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	dir, _ := os.Getwd()

	locationFile := filepath.Join(dir, constants.DIR_FILE_PERIZINAN, file_pendukung.Filename)
	dst, err := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	perizinan := model.Perizinan{
		Start:               start,
		Finish:              finish,
		FilePendukung:       pathFileInServer,
		KategoriPerizinanID: kategori_perizinan_id,
	}

	if !svc.SvcPegawai.UpdatePerizinan(perizinan_id, perizinan) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "maximum hari yang diperbolehkan mengajukan perizinan ini melebihi batas range izin",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "Pengajuan perizinan anda berhasil diperbaruhi, silahkan menunggu admin meng-acc perizinan anda",
	})

}

func (svc *PerizinanEchoController) GetAllController(c echo.Context) error {

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]
	claimsToken, _ := middleware.DecodeTokenPegawai(token)

	perizinan := svc.SvcPegawai.GetAllPerizinan(int(claimsToken["user_id"].(float64)))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"total":   len(perizinan),
		"data":    perizinan,
	})
}

func (svc *PerizinanEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameters id",
			"success": false,
		})
	}
	perizinan := svc.SvcPegawai.GetByIdPerizinan(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    perizinan,
		"success": true,
	})
}
