package main

import (
	"codechiev/utils"
	"log"
	"os"
	"strconv"
	"test-grpc/test/stream"
)

func main() {
	if len(os.Args) > 1 {
		num, err := strconv.Atoi(os.Args[1])
		utils.PanicIf(err)
		log.Println("Number", num)
		stream.RunClient(int32(num))()
	} else {
		log.Fatalln("Must input a number!")
	}
}
