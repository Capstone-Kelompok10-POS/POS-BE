package helpers

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(codeMember string) error {
	// Generate the barcode
	barcodeEncode, err := code128.Encode(codeMember)
	if err != nil {
		fmt.Println(err.Error())
	}
	barcode, err := barcode.Scale(barcodeEncode, 500, 150)
	if err != nil {
		fmt.Println(err.Error())
	}
	file, err := os.Create(codeMember+".png")
	if err != nil {
		fmt.Println(err.Error())
	}

	png.Encode(file, barcode)

	return nil
}