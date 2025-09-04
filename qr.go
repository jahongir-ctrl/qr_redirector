package main

import "github.com/skip2/go-qrcode"

func qr() {
	url := "http://localhost:8080/app" // ссылка на твой сервер
	qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode.png")
}
