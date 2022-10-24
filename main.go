package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please enter the excel file!")
	}

	filePath := os.Args[1]
	file := GetExcel(filePath)

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	format(file)
}
