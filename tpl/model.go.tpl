package model

import (
	{{if .HasTime}}"time"{{end}}
)

type {{ .Model }} struct {
{{range $index, $col := .Cols}}	{{printf "%s %s `gorm:\"column:%s\"`\n" $col.Model $col.ModelType $col.Name}}{{end}}}

func ({{ .Model }}) TableName() string {
  return "{{ .Name }}"
}