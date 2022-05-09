package routes

import (
	"mini-project-go/config"
	"mini-project-go/constants"
	"mini-project-go/controller"
	a "mini-project-go/controller/admin"
	"mini-project-go/repository"
	"mini-project-go/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RegisterAuthAPI(e *echo.Echo, db *gorm.DB, conf config.Config) {
	repo := repository.NewAuthRepository(db)
	svc := service.NewServiceAuth(repo, conf)
	cont := controller.AuthEchoController{
		SvcAuth: svc,
	}

	e.POST("/login", cont.LoginUserController)
	e.POST("/logout", cont.LogoutUserController)
}

func RegisterAdminAPI(e *echo.Echo, db *gorm.DB, conf config.Config) {
	repo := repository.NewAdminRepository(db)
	svc := service.NewServiceAdmin(repo, conf)

	contJamKerja := a.JamKerjaEchoController{
		SvcAdmin: svc,
	}
	contJamKerjaDetail := a.JamKerjaDetailEchoController{
		SvcAdmin: svc,
	}

	r := e.Group("/admin")
	r.Use(middleware.JWT([]byte(constants.SCREET_JWT_FOR_ADMIN)))
	r.POST("/jamkerja", contJamKerja.CreateController)
	r.GET("/jamkerja", contJamKerja.GetAllController)
	r.GET("/jamkerja/:id", contJamKerja.GetByIdController)
	r.PUT("/jamkerja/:id", contJamKerja.UpdateController)
	r.DELETE("/jamkerja/:id", contJamKerja.DeleteController)
	//-------------------------------------------------------
	r.POST("/jamkerja/detail", contJamKerjaDetail.CreateController)
	r.GET("/jamkerja/detail/:id", contJamKerjaDetail.GetByIdController)
}
