package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bernoussama/shortlink/internal/db"
)

func Redirect(d *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		short := r.URL.Path[1:]
		url, err := db.GetURL(d, short)
		if err != nil {
			log.Println(err)
			ShowIndex(w, r)
			return
		}
		log.Println(url)
		if url == "" {
			ShowIndex(w, r)
			return
		}
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
