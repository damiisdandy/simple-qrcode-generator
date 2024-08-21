package main

import (
	"flag"
	"fmt"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	inputString := flag.String("input", "", "String you'd like to convert to QR code")
	outputDir := flag.String("output", "", "Directory where you'd like to save the QR code")
	fileName := flag.String("filename", "qr.png", "File name of the QR code")
	flag.Parse()

	if *inputString == "" {
		flag.Usage()
		return
	}

	filePath := *outputDir + "/" + *fileName

	if *outputDir == "" {
		filePath = *fileName
	}

	fmt.Printf("Writing %s QR code to %s\n", *inputString, filePath)

	err := qrcode.WriteFile(*inputString, qrcode.Highest, 512, filePath)
	if err != nil {
		log.Fatalf("Error writing QR code: %s", err)
		return
	}

	fmt.Println("QR code written to", filePath)
}
