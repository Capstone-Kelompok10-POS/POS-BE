package helpers

import (
	"fmt"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

func GenerateBarcode(Code_Member string) error {
	// Generate the barcode
	barcodeEncode, err := code128.Encode(Code_Member)
	if err != nil {
		fmt.Println(err.Error())
	}
	barcode, err := barcode.Scale(barcodeEncode, 500, 150)
	if err != nil {
		fmt.Println(err.Error())
	}
	file, err := os.Create(Code_Member+".png")
	if err != nil {
		fmt.Println(err.Error())
	}

	png.Encode(file, barcode)

	return nil
}