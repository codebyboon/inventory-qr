package main

import (
	qr "github.com/codebyboon/inventory-qr/internal/infra/qrcode"
	logger "github.com/codebyboon/inventory-qr/pkg/logger"
)

func main() {

	data := "https://example.com"
	filePath := "./qr-output/qr.png"
	err := qr.GenerateQRCode(data, filePath)
	if err != nil {
		logger.ErrorLogger.Fatalf("Failed to generate QR code: %v", err)
	}

	logger.InfoLogger.Println("QR code generated successfully:", filePath)
}
