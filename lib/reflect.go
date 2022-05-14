package lib

import (
	"mini-project-go/constants"
	"net/mail"
	"reflect"
	"strings"
	"time"

	"github.com/jftuga/geodist"
	"golang.org/x/crypto/bcrypt"
)

func CheckDataType(m interface{}) bool {
	rt := reflect.TypeOf(m)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}

func InterfaceToSliceStr(m []interface{}) []string {
	var res []string
	for _, v := range m {
		res = append(res, v.(string))
	}
	return res
}

func ValidateDay(m []string) bool {
	dnc := len(m)
	cnt := 0
	for _, vd := range constants.DAY {
		for _, dn := range m {
			if vd == dn {
				cnt++
			}
		}
	}
	if dnc == cnt {
		return true
	} else {
		return false
	}
}

func ValidateDayJamKerjaDetail(m []map[string]interface{}) bool {

	day := []string{}

	for i := 0; i < len(m); i++ {
		day = append(day, m[i]["hari"].(string))
	}

	if !ValidateDay(day) {
		return false
	} else {
		return true
	}
}

func CreateSliceJamKerjaDetail(m []interface{}, jamkerja_id int) []map[string]interface{} {

	var res []map[string]interface{}

	for i, v := range m {
		res = append(res, v.(map[string]interface{}))
		res[i]["jamkerja_id"] = jamkerja_id
		res[i]["created_at"] = time.Now()
		res[i]["updated_at"] = time.Now()
	}

	return res

}

func MakePassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func EmailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func CheckExtensionForImage(extension string) bool {
	allowed := []string{
		".png", ".jpg", ".jpeg",
	}
	extension = strings.ToLower(extension)
	for _, v := range allowed {
		if v == extension {
			return true
		}
	}
	return false
}

func CheckDistance(latPegawaiFloat float64, longPegawaiFloat float64, latUnitKerja float64, lonUnitKerja float64) bool {
	var pegawaiLoc = geodist.Coord{
		Lat: latPegawaiFloat,
		Lon: longPegawaiFloat,
	}
	var unitKerjaLoc = geodist.Coord{
		Lat: latUnitKerja,
		Lon: lonUnitKerja,
	}
	miles, _, _ := geodist.VincentyDistance(pegawaiLoc, unitKerjaLoc)
	if miles <= constants.MIN_DISTANCE_ABSENSI {
		return true
	} else {
		return false
	}
}

// func InTimeSpan(start, end, check time.Time) bool {
// 	if start.Before(end) {
// 		return !check.Before(start) && !check.After(end)
// 	}
// 	if start.Equal(end) {
// 		return check.Equal(start)
// 	}
// 	return !start.After(check) || !end.Before(check)
// }
// func StrToTimeFormat(str string) time.Time {
// 	res, _ := time.Parse(constants.FORMAT_TIME, str)
// 	return res
// }

func Strtotime(str string) int64 {
	layout := "15:04:05"
	t, _ := time.Parse(layout, str)
	return t.Unix()
}
