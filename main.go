package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	ua := strings.ToLower(r.UserAgent())

	_, err := db.Exec("INSERT INTO scans (user_agent) VALUES ($1)", ua)
	if err != nil {
		log.Println("Ошибка записи в БД:", err)
	}

	if strings.Contains(ua, "android") {
		http.Redirect(w, r, "https://play.google.com/store/apps/details?id=tj.dc.myid1a&hl=ru", http.StatusFound)
		return
	}
	if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		http.Redirect(w, r, "https://apps.apple.com/tj/app/%D0%B8%D0%BC%D0%B7%D0%BE/id6748394310", http.StatusFound)
		return
	}
	http.Redirect(w, r, "https://play.google.com/store/apps/details?id=tj.dc.myid1a&hl=ru", http.StatusFound)

}

func StatsHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, err := db.Query("SELECT id, created_at, user_agent FROM scans ORDER BY created_at DESC LIMIT 20")
	if err != nil {
		http.Error(w, "Ошибка при получении статистики", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	fmt.Fprintln(w, "последние 20 переходов:\n")
	for rows.Next() {
		var id int
		var createdAt, ua string
		if err := rows.Scan(&id, &createdAt, &ua); err == nil {
			fmt.Fprintf(w, "%d | %s | %s\n", id, createdAt, ua)
		} else {
			fmt.Println("Ошибка при чтении строки:", err)
		}
	}
}

func main() {
	db := InitDB()
	defer db.Close()
	http.HandleFunc("/app", func(w http.ResponseWriter, r *http.Request) {
		RedirectHandler(w, r, db)
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		StatsHandler(w, r, db)
	})

	port := "8080"
	fmt.Println("Сервер запущен на порт", port)
	http.ListenAndServe(":"+port, nil)
}
