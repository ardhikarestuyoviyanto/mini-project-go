package admin

import (
	"mini-project-go/domain"
	"mini-project-go/lib"
	"mini-project-go/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PegawaiEchoController struct {
	SvcAdmin domain.AdapterAdminService
}

func (svc *PegawaiEchoController) CreateController(c echo.Context) error {
	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	nama, namaErr := dataRaw["nama"].(string)
	email, emailErr := dataRaw["email"].(string)
	password, passwordErr := dataRaw["password"].(string)
	jabatan, jabatanErr := dataRaw["jabatan"].(string)
	unitkerjaId, unitkerjaIdErr := dataRaw["unitkerja_id"].(float64)

	if !namaErr || !emailErr || !passwordErr || !jabatanErr || !unitkerjaIdErr || !lib.EmailValidation(email) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}

	if svc.SvcAdmin.GetByEmailPegawai(email).Nama != "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "email telah digunakan users lain",
		})
	}

	pegawai := model.Users{
		Nama:        nama,
		Email:       email,
		Password:    lib.MakePassword(password),
		Jabatan:     jabatan,
		UnitkerjaID: int(unitkerjaId),
		RoleID:      2,
	}

	svc.SvcAdmin.CreatePegawai(pegawai)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"success": true,
		"message": "akun pegawai baru berhasil disimpan",
	})
}

func (svc *PegawaiEchoController) GetAllController(c echo.Context) error {
	pegawai := svc.SvcAdmin.GetAllPegawai()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"total":   len(pegawai),
		"data":    pegawai,
	})
}

func (svc *PegawaiEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id parameters",
			"success": false,
		})
	}
	pegawai := svc.SvcAdmin.GetByIdPegawai(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data":    pegawai,
	})
}

func (svc *PegawaiEchoController) UpdateController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}
	nama, namaErr := dataRaw["nama"].(string)
	email, emailErr := dataRaw["email"].(string)
	password, passwordErr := dataRaw["password"].(string)
	jabatan, jabatanErr := dataRaw["jabatan"].(string)
	unitkerjaId, unitkerjaIdErr := dataRaw["unitkerja_id"].(float64)

	if !namaErr || !emailErr || !passwordErr || !jabatanErr || !unitkerjaIdErr || !lib.EmailValidation(email) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": "invalid form input",
		})
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid parameters id",
			"success": false,
		})
	}

	if len(password) == 0 {
		pegawai := model.Users{
			Nama:        nama,
			Email:       email,
			Jabatan:     jabatan,
			UnitkerjaID: int(unitkerjaId),
		}
		svc.SvcAdmin.UpdatePegawai(id, pegawai)
	} else {
		pegawai := model.Users{
			Nama:        nama,
			Email:       email,
			Password:    lib.MakePassword(password),
			Jabatan:     jabatan,
			UnitkerjaID: int(unitkerjaId),
		}
		svc.SvcAdmin.UpdatePegawai(id, pegawai)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "akun pegawai berhasil diupdate",
	})

}

func (svc *PegawaiEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id parameters",
			"status":  false,
		})
	}
	svc.SvcAdmin.DeletePegawai(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "akun pegawai berhasil dihapus",
		"status":  true,
	})
}
