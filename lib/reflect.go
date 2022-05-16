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

func CheckExtensionForPdf(extension string) bool {
	if strings.ToLower(extension) == ".pdf" {
		return true
	} else {
		return false
	}
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
	if (miles * constants.MILES_TO_METERS) <= constants.MIN_DISTANCE_ABSENSI {
		return true
	} else {
		return false
	}
}

func Strtotime(str string) int64 {
	layout := "15:04:05"
	t, _ := time.Parse(layout, str)
	return t.Unix()
}

func RangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func ValidationDatePerizinan(starts, finishs string, max_izin int) bool {

	dateStarts, errDateStarts := time.Parse("2006-01-02", starts)
	dateFinishs, errDateFinish := time.Parse("2006-01-02", finishs)

	start := time.Date(dateStarts.Year(), dateStarts.Month(), dateStarts.Day(), 0, 0, 0, 0, dateStarts.Local().Location())
	end := start.AddDate(0, 0, (dateFinishs.Day() - dateStarts.Day()))

	counter := 0

	if dateStarts.Month() != dateFinishs.Month() || errDateFinish != nil || errDateStarts != nil {
		return false
	}

	for rd := RangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		counter++
		// fmt.Println(date.Format("2006-01-02"))
	}
	if counter > max_izin {
		return false
	} else {
		return true
	}
}
