package main

import (
	"errors"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/liyue201/goqr"
)

func scanQRCode(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not open file: %v\n", err))
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Could not decode image: %v\n", err))
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Recognize failed: %v\n", err))
	}

	return fmt.Sprintf("QR Code Content: %s", qrCodes[0].Payload), nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No File given.")
	}

	file := os.Args[1]
	qrContent, _ := scanQRCode(file)

	fmt.Println(qrContent)
}
