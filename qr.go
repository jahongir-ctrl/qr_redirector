package main

import "github.com/skip2/go-qrcode"

func main() {
	url := "https://qr-redirector-1.onrender.com/app"
	qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode.png")
}
