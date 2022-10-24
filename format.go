package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Table struct {
	Name   string
	Fields []interface{}
}

func format(f *excelize.File) {

	sheets := f.GetSheetList()

	tables := []*Table{}

	for _, sheet := range sheets {
		fmt.Println(sheet)

		// 获取 Sheet1 上所有单元格
		rows, err := f.GetRows(sheet)
		if err != nil {
			panic(err.Error())
		}

		// 字段属性名称
		var cols []string
		if len(rows) > 3 {
			for _, col := range rows[2] {
				cols = append(cols, strings.ReplaceAll(col, " ", "_"))
			}
		} else {
			panic(errors.New("cols is null"))
		}
		colsLen := len(cols)

		// 所有字段
		fields := []interface{}{}
		for i, row := range rows {
			// 第3行往后
			if i > 3 {
				field := map[string]string{}
				// 每列属性
				for j, colValue := range row {
					if j < colsLen {
						colName := cols[j]
						if j == 0 {
							colName = "ID"
						}
						field[colName] = colValue
					}
				}
				fields = append(fields, field)
			}
		}

		table := Table{Name: sheet, Fields: fields}
		tables = append(tables, &table)
	}

	data, err := json.MarshalIndent(tables, "", "  ")
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(data))

	os.WriteFile("data.json", data, 0666)
}
