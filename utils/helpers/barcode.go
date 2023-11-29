package helpers

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(CodeMember string) error {
	// Generate the barcode
	barcodeEncode, err := code128.Encode(CodeMember)
	if err != nil {
		fmt.Println(err.Error())
	}
	barcode, err := barcode.Scale(barcodeEncode, 500, 150)
	if err != nil {
		fmt.Println(err.Error())
	}
	file, err := os.Create(CodeMember+".png")
	if err != nil {
		fmt.Println(err.Error())
	}

	png.Encode(file, barcode)

	return nil
}