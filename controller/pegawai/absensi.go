package pegawai

import (
	"io"
	"mini-project-go/constants"
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/middleware"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
)

type PegawaiEchoController struct {
	SvcPegawai domain.AdapterPegawaiService
}

func (svc *PegawaiEchoController) AbsenMasukController(c echo.Context) error {

	masuk := c.FormValue("masuk")
	foto, err := c.FormFile("foto")
	latPegawai := c.FormValue("lat")
	longPegawai := c.FormValue("long")

	latPegawaiFloat, errLatFloat := strconv.ParseFloat(latPegawai, 64)
	longPegawaiFloat, errLongFloat := strconv.ParseFloat(longPegawai, 64)

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]
	claimsToken, _ := middleware.DecodeTokenPegawai(token)

	pathImgInServer := os.Getenv("BASE_URL") + constants.STATIC_FILE_FOTO_ABSEN + foto.Filename

	if err != nil || len(masuk) == 0 || errLatFloat != nil || errLongFloat != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid form",
		})
	}

	if !lib.CheckExtensionForImage(filepath.Ext(pathImgInServer)) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid extension file",
		})
	}

	carbon := carbon.SetLanguage(lib.CarbonLanguageNew())
	today := carbon.Now().ToDateString()

	mMasuk, bMasuk, _, _ := svc.SvcPegawai.GetJamKerjaDetailTodayByIdUnitKerja(int(claimsToken["unitkerja_id"].(float64)), carbon.Now().ToWeekString())
	latUnitKerja, lonUnitKerja := svc.SvcPegawai.GetLatiduteLongtiduteUnitKerja(int(claimsToken["unitkerja_id"].(float64)))

	if !svc.SvcPegawai.CheckHariLibur(int(claimsToken["unitkerja_id"].(float64)), carbon.Now().ToWeekString()) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Sekarang adalah hari libur pegawai",
		})
	}

	if lib.Strtotime(masuk) >= lib.Strtotime(mMasuk) && lib.Strtotime(masuk) <= lib.Strtotime(bMasuk) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Bukan waktunya absen masuk ",
		})
	}

	if !lib.CheckDistance(latPegawaiFloat, longPegawaiFloat, latUnitKerja, lonUnitKerja) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Jarak lokasi anda dan lokasi unit kerja terlalu jauh",
		})
	}

	masukAbsn, _ := svc.SvcPegawai.GetPresensiToday(int(claimsToken["user_id"].(float64)), today)

	if masukAbsn != "00:00:00" && len(masukAbsn) != 0 {
		if !lib.CheckDistance(latPegawaiFloat, longPegawaiFloat, latUnitKerja, lonUnitKerja) {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"success": false,
				"message": "Anda telah melakukan presensi masuk hari ini",
			})
		}
	}

	//-----------------------------------------------------------------------------

	src, err := foto.Open()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	dir, _ := os.Getwd()

	locationFile := filepath.Join(dir, constants.DIR_FILE_FOTO_ABSEN, foto.Filename)
	dst, err := os.OpenFile(locationFile, os.O_WRONLY|os.O_CREATE, 06666)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if len(masukAbsn) == 0 {
		svc.SvcPegawai.ActionInsertAbsenMasuk(int(claimsToken["user_id"].(float64)), today, masuk, pathImgInServer)
	} else {
		svc.SvcPegawai.ActionUpdateAbsenMasuk(int(claimsToken["user_id"].(float64)), today, masuk, pathImgInServer)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Anda berhasil melakukan presensi masuk",
		"success": true,
	})
}
