package service_qrcode

import (
	"bufio"
	"image"
	"log"
	"net/http"
	"os"
)

func ReadLinks(file string) *[]string {
	pfile, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer pfile.Close()

	scanner := bufio.NewScanner(pfile)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var links []string
	for scanner.Scan() {
		links = append(links, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &links
}

func GetImage(imgUrl string) *image.Image {
	resp, err := http.Get(imgUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return &img
}
