package qr

import (
	"os"
	"path/filepath"

	logger "github.com/codebyboon/inventory-qr/pkg/logger"
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(data string, filePath string) error {

	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			logger.ErrorLogger.Printf("Error creating directory: %v", err)
			return err
		}
	}
	return qrcode.WriteFile(data, qrcode.Medium, 256, filePath)
}
