package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/index.html", http.StatusSeeOther)

	/*
		tmpl, err := template.ParseFiles("static/index.html")
		if err != nil {
			http.Error(w, "Error loading page", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("wtf %v", err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	*/
}
