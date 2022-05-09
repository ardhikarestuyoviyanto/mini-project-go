package lib

import (
	"mini-project-go/constants"
	"reflect"
	"time"
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
