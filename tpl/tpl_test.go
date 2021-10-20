package tpl

import (
	"fmt"
	mssqlCfg "sd-2110/config/mssql"
	"sd-2110/tpl/model"
	"testing"
)

func TestTpl(t *testing.T) {
	cols := []Col{{Name: "col1", Type: "string"}, {Name: "col2", Type: "datetimeoffset"}}
	table := Table{Name: "users", Cols: cols}

	TableToModel(&table)
}

func TestModel(t *testing.T) {
	db := mssqlCfg.GetDB()
	var users = []model.Users1{}
	db.Debug().Find(&users)
	fmt.Println(users)
}
