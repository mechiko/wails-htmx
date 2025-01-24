package domain

import "strings"

var CisTypes = map[string]string{
	"emitted":          "эмитирован",
	"applied":          "нанесен",
	"introduced":       "в обороте",
	"written_off":      "списан",
	"retired":          "выбыл",
	"withdrow":         "выбыл",
	"disaggregation":   "расформирован",
	"disaggregated":    "расформирован",
	"applied_not_paid": "неоплачен",
}

func DictCisTypes(s string) string {
	if s == "" {
		return ""
	}
	if out, ok := CisTypes[strings.ToLower(s)]; ok {
		return out
	}
	return s
}
