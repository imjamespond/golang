package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestExcelSheets(t *testing.T) {

	f := getExcel()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	format(f)
}

func TestExcel(t *testing.T) {

	f := getExcel()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 获取工作表中指定单元格的值
	cell, err := f.GetCellValue("当事人T_PARTY", "A3")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cell)

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("当事人T_PARTY")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func getExcel() *excelize.File {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	filePath := fmt.Sprintf("%s/Downloads/模型演示数据.xlsx", home)

	return GetExcel(filePath)
}
