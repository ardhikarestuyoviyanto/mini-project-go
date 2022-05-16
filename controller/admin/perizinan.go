package admin

import (
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type PerizinanEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *PerizinanEchoController) GetAllController(c echo.Context) error {
	perizinan := svc.SvcAdmin.GetAllPerizinan()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    perizinan,
		"total":   len(perizinan),
	})
}

func (svc *PerizinanEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "Invalid parameters id",
		})
	}
	perizinan := svc.SvcAdmin.GetByIdPerizinan(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    perizinan,
		"success": true,
	})
}

func (svc *PerizinanEchoController) UpdateController(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameters id",
			"success": "false",
		})
	}

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	catatan, catatanErr := dataRaw["catatan"].(string)
	status, statusErr := dataRaw["status"].(string)

	if !catatanErr || !statusErr || len(catatan) == 0 || len(status) == 0 || status != "N" && status != "Y" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid form input",
			"success": false,
		})
	}

	if status == "N" {
		perizinan := model.Perizinan{
			Catatan: catatan,
			Status:  status,
		}
		svc.SvcAdmin.UpdatePerizinan(id, perizinan)
	} else {
		perizinan := model.Perizinan{
			Catatan: catatan,
			Status:  status,
		}
		svc.SvcAdmin.UpdatePerizinan(id, perizinan)

		perizinanDetail := svc.SvcAdmin.GetByIdPerizinan(id)
		svc.SvcAdmin.DeleteRekapAbsen(perizinanDetail.Start, perizinanDetail.Finish, perizinanDetail.UserId)
		dateStarts, _ := time.Parse("2006-01-02", perizinanDetail.Start)
		dateFinishs, _ := time.Parse("2006-01-02", perizinanDetail.Finish)

		start := time.Date(dateStarts.Year(), dateStarts.Month(), dateStarts.Day(), 0, 0, 0, 0, dateStarts.Local().Location())
		end := start.AddDate(0, 0, (dateFinishs.Day() - dateStarts.Day()))

		for rd := lib.RangeDate(start, end); ; {
			date := rd()
			if date.IsZero() {
				break
			}
			rekapAbsenNew := model.RekapAbsen{
				Tanggal:     date.Format("2006-01-02"),
				Keterangan:  "Pegawai Sedang Izin",
				UserID:      perizinanDetail.UserId,
				PerizinanID: perizinanDetail.ID,
				Masuk:       "00:00:00",
				Pulang:      "00:00:00",
			}
			svc.SvcAdmin.InsertRekapAbsen(rekapAbsenNew)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pengajuan perizinan pegawai berhasil diproses",
		"success": true,
	})
}

func (svc *PerizinanEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameters id",
			"success": "false",
		})
	}
	perizinanDetail := svc.SvcAdmin.GetByIdPerizinan(id)
	rekapabsen := model.RekapAbsen{
		Keterangan: "TIDAK HADIR",
		Masuk:      "00:00:00",
		Pulang:     "00:00:00",
	}
	svc.SvcAdmin.UpdateRekapAbsen(perizinanDetail.Start, perizinanDetail.Finish, perizinanDetail.UserId, rekapabsen)
	svc.SvcAdmin.DeletePerizinan(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "perizinan pegawai berhasil dihapus",
		"success": true,
	})
}
