package tpl

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"my.com/utils"
)

type Col struct {
	Name string
	Type string
}
type Table struct {
	Name string
	Cols []Col
}

func (c Col) Model() string {
	return strings.Title(c.Name)
}
func (c Col) ModelType() string {
	return TypeMssqlDict[c.Type]
}
func (t Table) Model() string {
	return strings.Title(t.Name)
}
func (t Table) HasTime() bool {
	for _, col := range t.Cols {
		if col.Type == "datetimeoffset" {
			return true
		}
	}
	return false
}

// https://juejin.cn/post/6844903762901860360
// https://studygolang.com/static/pkgdoc/pkg/text_template.htm
func TableToModel(table *Table) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	// wd, err := os.Getwd()
	// utils.FatalIf(err)
	// fmt.Println(basepath)

	tpl, err := template.ParseFiles(basepath + "/model.go.tpl")
	utils.PanicIf(err)

	file, err := os.OpenFile(basepath+"/.output/"+table.Name+".go", os.O_WRONLY|os.O_CREATE, 0644)
	utils.PanicIf(err)
	file.Truncate(0)
	file.Seek(0, 0)
	err = tpl.Execute(
		file, // os.Stdout,
		table,
	)
	utils.PanicIf(err)
}
