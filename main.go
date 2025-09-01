package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
	file    = "counter.txt"
)

func loadCounter() {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		counter = 0
		return
	}
	counter, _ = strconv.Atoi(strings.TrimSpace(string(data)))
}

func saveCounter() {
	ioutil.WriteFile(file, []byte(strconv.Itoa(counter)), 0644)
}

func main() {
	loadCounter()
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		counter++
		saveCounter()
		mu.Unlock()

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
		fmt.Printf("всего отсканировано : %d", counter)
	})

	port := "8080"
	fmt.Println("Сервер запущен на порт", port)
	http.ListenAndServe(":"+port, nil)
}
