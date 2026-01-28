package main

import (
	"fmt"
)

func main() {
	ip := GetLocalIP()

	fmt.Printf("Local IP is %s\n", ip)

	url := fmt.Sprintf("http://%s:%d", ip, 8080)

	fmt.Printf("Server URL: %s\n", url)

	qrBase64, err := GenerateQR(url)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("QR (base64): %s\n", qrBase64)
}
