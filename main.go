package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/signal"
	"syscall"

	"felarof/core"
)

//go:embed assets/index.html
var assetsFS embed.FS

func main() {
	tmpl, err := template.ParseFS(assetsFS, "assets/index.html")
	if err != nil {
		log.Fatalf("Failed to load template: %v", err)
	}

	files := os.Args[1:]

	server, err := core.NewServer(files, tmpl)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer server.Close()

	url := server.GetURL()
	fmt.Printf("Server URL: %s\n", url)

	qrBase64, err := core.GenerateQR(url)
	if err != nil {
		log.Fatalf("Failed to generate QR: %v", err)
	}
	fmt.Printf("QR (base64): %s\n", qrBase64)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
