package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
)

var counter uint64

func main() {
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&counter, 1)
		ua := strings.ToLower(r.UserAgent())

		if strings.Contains(ua, "android") {
			http.Redirect(w, r, "https://play.google.com/store/apps/details?id=tj.dc.myid1a&hl=ru", http.StatusFound)
			return
		}
		if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
			http.Redirect(w, r, "https://apps.apple.com/tj/app/%D0%B8%D0%BC%D0%B7%D0%BE/id6748394310", http.StatusFound)
			return
		}
		http.Redirect(w, r, "https://play.google.com/store/apps/details?id=tj.dc.myid1a&hl=ru", http.StatusFound)

	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("всего отсканировано : %d", atomic.LoadUint64(&counter))
	})

	port := "8080"
	fmt.Println("Сервер запущен на порт", port)
	http.ListenAndServe(":"+port, nil)
}
