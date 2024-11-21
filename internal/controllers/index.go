package controllers

import (
	"fmt"
	"net/http"
	"text/template"
)

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	viewsDir := "internal/views"
	templ, err := template.ParseFiles(fmt.Sprintf("%s/index.html", viewsDir))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = templ.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
