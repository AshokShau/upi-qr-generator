package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/AshokShau/upi-qr-generator/str"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.New("index").Parse(str.Index)
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("wtf %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}

}
