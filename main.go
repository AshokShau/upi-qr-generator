package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshokShau/upi-qr-generator/str"
)

func main() {
	http.HandleFunc("/", str.Home)
	http.HandleFunc("/qr", str.GenerateQRCodeHandler)
	http.HandleFunc("/pay", str.RenderPaymentPageHandler)

	fmt.Println("Server started on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
