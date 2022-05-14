package controller

import (
	"mini-project-go/constants"
	"mini-project-go/domain"
	"mini-project-go/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthEchoController struct {
	SvcAuth domain.AdapterAuthService
}

func (svc *AuthEchoController) LoginUserController(c echo.Context) error {

	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return err
	}

	email, errEmail := dataRaw["email"].(string)
	password, errPassword := dataRaw["password"].(string)

	if !errEmail || !errPassword {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": "Masukkan email dan password",
		})
	}

	res, err := svc.SvcAuth.LoginUsers(email, password)

	if !err {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"success": false,
			"message": "Akun tidak ditemukan",
		})
	} else {

		if res.RoleID == 1 {
			token, _ := middleware.CreateTokenAdmin(res.ID, res.RoleID, res.Email, res.Nama)
			svc.SvcAuth.UpdateToken(token, int(res.ID))
			return c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"role":    "Admin",
				"token":   token,
			})
		} else {
			token, _ := middleware.CreateTokenPegawai(res.ID, res.RoleID, res.UnitkerjaID, res.Email, res.Nama)
			svc.SvcAuth.UpdateToken(token, int(res.ID))
			return c.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"role":    "Pegawai",
				"token":   token,
			})
		}

	}
}

func (svc *AuthEchoController) LogoutUserController(c echo.Context) error {

	tokenHeader := c.Request().Header.Get("Authorization")
	token := tokenHeader[len(constants.TOKEN_JWT_TYPE):]

	if len(token) == 0 {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	res := svc.SvcAuth.GetUsersByToken(token)

	if res.Token == "" {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}

	svc.SvcAuth.LogoutUsers(res.ID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "User successfully logout",
	})

}
