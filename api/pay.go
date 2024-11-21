package handler

import (
	"github.com/AshokShau/upi-qr-generator/str"
	"net/http"
)

func Pay(w http.ResponseWriter, r *http.Request) {
	str.RenderPaymentPageHandler(w, r)
}
