package main

import (
	"fmt"
	"log"
)

func main() {
	ip := GetLocalIP()
	fmt.Printf("Local IP is %s\n", ip)

	token, err := GenerateToken()
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	url := fmt.Sprintf("http://%s:%d/%s", ip, 8080, token)
	fmt.Printf("Server URL: %s\n", url)

	qrBase64, err := GenerateQR(url)
	if err != nil {
		log.Fatalf("Failed to generate QR code: %v", err)
	}

	fmt.Printf("QR (base64): %s\n", qrBase64)
}
