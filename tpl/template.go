package tpl

import (
	"io"
	"text/template"
)

var funcs = template.FuncMap{
	"mult": func(n1 float64, n2 float64) float64 {
		return n1 * n2
	},
	"sum": func(n1 float64, n2 float64) float64 {
		return n1 + n2
	},
	"sub": func(n1 float64, n2 float64) float64 {
		return n1 - n2
	},
}

func Execute(content string, data map[string]any, w io.Writer) error {
	tpl, err := template.New("tpl").Funcs(funcs).Parse(content)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}
