package model

type QRCodeConfig struct {
	Width  int
	Heigth int
	CodeX  int
	CodeY  int
}

func GetQRCodeConfig(cfg map[string]interface{}) *QRCodeConfig {
	codeX := int(cfg["x"].(float64))
	codeY := int(cfg["y"].(float64))
	width := int(cfg["width"].(float64))
	heigth := int(cfg["heigth"].(float64))

	return &QRCodeConfig{Width: width, Heigth: heigth, CodeX: codeX, CodeY: codeY}
}
