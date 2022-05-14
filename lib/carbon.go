package lib

import "github.com/golang-module/carbon/v2"

func CarbonLanguageNew() *carbon.Language {
	lang := carbon.NewLanguage()
	resources := map[string]string{
		"months":         "january|february|march|april|may|june|july|august|september|october|november|december",
		"short_months":   "jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec",
		"weeks":          "minggu|senin|selasa|rabu|kamis|jumat|sabtu",
		"short_weeks":    "sun|mon|tue|wed|thu|fri|sat",
		"seasons":        "spring|summer|autumn|winter",
		"constellations": "aries|taurus|gemini|cancer|leo|virgo|libra|scorpio|sagittarius|capricornus|aquarius|pisce",
		"year":           "1 yr|%d yrs",
		"month":          "1 mo|%d mos",
		"week":           "%dw",
		"day":            "%dd",
		"hour":           "%dh",
		"minute":         "%dm",
		"second":         "%ds",
		"now":            "just now",
		"ago":            "%s ago",
		"from_now":       "in %s",
		"before":         "%s before",
		"after":          "%s after",
	}
	lang.SetResources(resources)
	return lang
}
