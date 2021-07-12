package service_qrcode

import (
	"image"
	"image/draw"
	_ "image/jpeg" // decode
	"image/png"
	"log"
	"os"
	"path/filepath"

	"4d-qrcode/model"
	"4d-qrcode/util"

	"github.com/nfnt/resize"
)

func OpenJPEG(file string) *image.Image {
	// Decode the JPEG data. If reading from file, create a reader with
	reader, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return &m
}

type process_func func(_tpl *image.Image, _code *image.Image, codeName string)

func Process(outputDir string, cfg *model.QRCodeConfig) process_func {
	return func(_tpl *image.Image, _code *image.Image, codeName string) {
		code := resize.Resize(uint(cfg.Width), uint(cfg.Heigth), *_code, resize.Lanczos3)
		// tpl := resize.Resize(150, 0, *_tpl, resize.Lanczos3) //preserve aspect ratio
		tpl := *_tpl

		img := image.NewRGBA(tpl.Bounds())
		draw.Draw(img, img.Bounds(), tpl, image.Point{0, 0}, draw.Src)
		draw.Draw(img, img.Bounds(), code, image.Point{cfg.CodeX, cfg.CodeY}, draw.Src)

		toimg, _ := os.Create(filepath.Join(outputDir, codeName))
		defer toimg.Close()

		// jpeg.Encode(toimg, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		enc := &png.Encoder{
			CompressionLevel: png.BestSpeed,
		}
		err := enc.Encode(toimg, img)
		util.FatalIf(err)
	}
}
