package middleware

import (
	"mini-project-go/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateTokenAdmin(id int, role_id int, email string, nama string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["role_id"] = role_id
	claims["email"] = email
	claims["nama"] = nama
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SCREET_JWT_FOR_ADMIN))
}

func CreateTokenPegawai(id int, role_id int, unitkerja_id int, email string, nama string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id
	claims["unitkerja_id"] = unitkerja_id
	claims["role_id"] = role_id
	claims["email"] = email
	claims["nama"] = nama
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SCREET_JWT_FOR_PEGAWAI))
}
