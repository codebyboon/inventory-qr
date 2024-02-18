package main

import (
	"github.com/codebyboon/inventory-qr/internal/infra/qrcode/"
)

func main() {
	err := qrcode.WriteFile("https://example.com", qrcode.Medium, 256, "qr.png")
	if err != nil {
	}
}
