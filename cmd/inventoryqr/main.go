package main

import (
	"log"

	qr "github.com/codebyboon/inventory-qr/internal/infra/qrcode"
)

func main() {
	err := qr.GenerateQRCode("https://example.com", "qr.png")
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	log.Println("QR code generated successfully:", "./")
}
