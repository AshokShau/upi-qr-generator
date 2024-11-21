package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshokShau/upi-qr-generator/str"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", str.Home)
	r.HandleFunc("/qr", str.GenerateQRCodeHandler).Methods("GET")
	r.HandleFunc("/pay", str.RenderPaymentPageHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	fmt.Println("Server started at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
