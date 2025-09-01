package main

import "github.com/skip2/go-qrcode"

func main() {
	// Ссылка внутри QR-кода
	url := "http://localhost:8080/app"

	// Генерируем PNG (размер 256px)
	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qrcode.png")
	if err != nil {
		panic(err)
	}
}
