package database

import (
	"fmt"
	"mini-project-go/config"
	"mini-project-go/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf config.Config) *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString))

	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(
		&model.Role{},
		&model.JamKerja{},
		&model.JamKerjaDetail{},
		&model.UnitKerja{},
		&model.Users{},
		&model.KategoriPerizinan{},
		&model.Perizinan{},
		&model.RekapAbsen{},
	)

	return DB

}

func SeedDB(db *gorm.DB, conf config.Config) {
	CreateRoleSeed(db, conf)
	CreateAdminSeed(db, conf)
}
