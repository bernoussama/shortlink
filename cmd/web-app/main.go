package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bernoussama/shortlink/internal/controllers"
	"github.com/bernoussama/shortlink/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "shortlink.db")
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.Close()

	db.CreateTable(sqlite)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			controllers.Redirect(sqlite)(w, r)
			return
		}
		controllers.ShowIndex(w, r)
	})
	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	http.HandleFunc("GET /{$}", controllers.Redirect(sqlite))
	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
