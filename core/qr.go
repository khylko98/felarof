package main

import (
	"encoding/base64"
	"github.com/skip2/go-qrcode"
)

func GenerateQR(url string) (string, error) {
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(png), nil
}
