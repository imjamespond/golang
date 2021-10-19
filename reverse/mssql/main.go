package main

import (
	"fmt"
	mssqlCfg "sd-2110/config/mssql"
	"sd-2110/tpl"
)

type Col = tpl.Col
type Table = tpl.Table

var TableToModel func(t *Table) = tpl.TableToModel

func main() {
	db := mssqlCfg.GetDB()
	var tables []map[string]interface{}
	db.Table("information_schema.tables").Find(&tables)
	for _, table := range tables {
		fmt.Println(table)
		db.Table("information_schema.tables").Find(&tables)

		tableName := table["TABLE_NAME"].(string)
		mCols := make([]Col, 0)
		var cols []map[string]interface{}
		db.Table("information_schema.columns").Where("TABLE_NAME = ?", tableName).Find(&cols)
		for _, col := range cols {
			fmt.Println(
				col["TABLE_CATALOG"], col["TABLE_SCHEMA"], col["TABLE_NAME"], col["COLUMN_NAME"],
				col["ORDINAL_POSITION"], col["DATETIME_PRECISION"], col["DATA_TYPE"], col["IS_NULLABLE"],
			)

			mCols = append(mCols, Col{Name: col["COLUMN_NAME"].(string), Type: col["DATA_TYPE"].(string)})
		}

		mTable := Table{Name: tableName, Cols: mCols}
		TableToModel(&mTable)
	}

}
