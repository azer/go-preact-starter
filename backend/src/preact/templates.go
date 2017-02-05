package preact

import (
	"html/template"
	"strings"
)

var templates *template.Template

var funcMap = template.FuncMap{
	"Upper": func(s string) string {
		return strings.ToUpper(s)
	},
	"Lower": func(s string) string {
		return strings.ToLower(s)
	},
}

func init() {
	ReloadTemplates()
}

func ReloadTemplates() {
	var err error
	templates, err = template.New("").Funcs(funcMap).ParseGlob("./backend/templates/*")
	if err != nil {
		log.Info("Failed to load templates. Error: %v", err)
		panic(err)
	}
}
