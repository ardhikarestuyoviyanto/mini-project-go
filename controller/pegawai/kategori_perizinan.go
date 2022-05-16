package pegawai

import (
	"mini-project-go/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KategoriPerizinanEchoController struct {
	SvcPegawai domain.AdapterPegawaiService
}

func (svc *KategoriPerizinanEchoController) GetAllController(c echo.Context) error {
	perizinan := svc.SvcPegawai.GetAllKategoriPerizinan()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    perizinan,
		"total":   len(perizinan),
		"success": true,
	})
}
