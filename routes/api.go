package routes

import (
	"mini-project-go/config"
	"mini-project-go/constants"
	"mini-project-go/controller"
	a "mini-project-go/controller/admin"
	p "mini-project-go/controller/pegawai"
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
	contUnitKerja := a.UnitKerjaEchoController{
		SvcAdmin: svc,
	}
	contPegawai := a.PegawaiEchoController{
		SvcAdmin: svc,
	}
	contKategoriPerizinan := a.KategoriPerizinanEchoController{
		SvcAdmin: svc,
	}
	contPerizinan := a.PerizinanEchoController{
		SvcAdmin: svc,
	}

	r := e.Group("/admin")
	r.Use(middleware.JWT([]byte(constants.SCREET_JWT_FOR_ADMIN)))
	//--------------------------------------------------------
	r.POST("/jamkerja", contJamKerja.CreateController)
	r.GET("/jamkerja", contJamKerja.GetAllController)
	r.GET("/jamkerja/:id", contJamKerja.GetByIdController)
	r.PUT("/jamkerja/:id", contJamKerja.UpdateController)
	r.DELETE("/jamkerja/:id", contJamKerja.DeleteController)
	//-------------------------------------------------------
	r.POST("/jamkerja/detail", contJamKerjaDetail.CreateController)
	r.GET("/jamkerja/detail/:id", contJamKerjaDetail.GetByIdController)
	//-------------------------------------------------------------
	r.POST("/unitkerja", contUnitKerja.CreateController)
	r.GET("/unitkerja", contUnitKerja.GetAllController)
	r.GET("/unitkerja/:id", contUnitKerja.GetByIdController)
	r.PUT("/unitkerja/:id", contUnitKerja.UpdateController)
	r.DELETE("/unitkerja/:id", contUnitKerja.DeleteController)
	//----------------------------------------------------------------
	r.POST("/pegawai", contPegawai.CreateController)
	r.PUT("/pegawai/:id", contPegawai.UpdateController)
	r.GET("/pegawai/:id", contPegawai.GetByIdController)
	r.GET("/pegawai", contPegawai.GetAllController)
	r.DELETE("/pegawai/:id", contPegawai.DeleteController)
	//-----------------------------------------------------------------
	r.POST("/perizinan/kategori", contKategoriPerizinan.CreateController)
	r.GET("/perizinan/kategori", contKategoriPerizinan.GetAllController)
	r.GET("/perizinan/kategori/:id", contKategoriPerizinan.GetByIdController)
	r.PUT("/perizinan/kategori/:id", contKategoriPerizinan.UpdateController)
	r.DELETE("/perizinan/kategori/:id", contKategoriPerizinan.DeleteController)
	//---------------------------------------------------------------------
	r.GET("/perizinan", contPerizinan.GetAllController)
	r.GET("/perizinan/:id", contPerizinan.GetByIdController)
	r.PUT("/perizinan/:id", contPerizinan.UpdateController)
	r.DELETE("/perizinan/:id", contPerizinan.DeleteController)
}

func RegisterPegawaiAPI(e *echo.Echo, db *gorm.DB, conf config.Config) {
	repo := repository.NewPegawaiRepository(db)
	svc := service.NewServicePegawai(repo, conf)
	contAbsensi := p.PegawaiEchoController{
		SvcPegawai: svc,
	}
	contPerizinan := p.PerizinanEchoController{
		SvcPegawai: svc,
	}
	contKategoriPerizinan := p.KategoriPerizinanEchoController{
		SvcPegawai: svc,
	}
	p := e.Group("/pegawai")
	p.Use(middleware.JWT([]byte(constants.SCREET_JWT_FOR_PEGAWAI)))
	//----------------------------------------------------------------------------------
	p.POST("/absen/masuk", contAbsensi.AbsenMasukController)
	p.POST("/absen/pulang", contAbsensi.AbsenPulangController)
	//------------------------------------------------------------------------------------
	p.POST("/perizinan", contPerizinan.CreateController)
	p.GET("/perizinan", contPerizinan.GetAllController)
	p.PUT("/perizinan/:id", contPerizinan.UpdateController)
	p.GET("/perizinan/:id", contPerizinan.GetByIdController)
	//------------------------------------------------------------------------------------
	p.GET("/perizinan/kategori", contKategoriPerizinan.GetAllController)
}
