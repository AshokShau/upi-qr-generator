package str

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/skip2/go-qrcode"
)

func Home(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("wtf %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// GenerateUPIURL creates the UPI URL string based on the request data.
func GenerateUPIURL(vpa, name string, amount float64) string {
	return fmt.Sprintf("upi://pay?pa=%s&pn=%s&am=%0.2f", vpa, name, amount)
}

// GenerateQRCodeHandler generates and serves the QR code as an image.
func GenerateQRCodeHandler(w http.ResponseWriter, r *http.Request) {
	vpa := r.URL.Query().Get("vpa")
	name := r.URL.Query().Get("name")
	amountStr := r.URL.Query().Get("amount")

	if vpa == "" || amountStr == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	if name == "" {
		name = "Unknown"
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	upiURL := GenerateUPIURL(vpa, name, amount)
	qr, err := qrcode.New(upiURL, qrcode.Medium)
	if err != nil {
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	err = qr.Write(256, w)
	if err != nil {
		http.Error(w, "Failed to write QR code", http.StatusInternalServerError)
		return
	}
}

// RenderPaymentPageHandler renders the payment gateway page.
func RenderPaymentPageHandler(w http.ResponseWriter, r *http.Request) {
	vpa := r.URL.Query().Get("vpa")
	name := r.URL.Query().Get("name")
	amount := r.URL.Query().Get("amount")

	// Load HTML template
	tmpl, err := template.ParseFiles("static/template.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	if name == "" {
		name = "Unknown"
	}

	if vpa == "" || amount == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	// Render template with data
	data := map[string]string{
		"VPA":    vpa,
		"Name":   name,
		"Amount": amount,
	}

	// Execute template with the data
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
