package handler

import (
	"github.com/AshokShau/upi-qr-generator/str"
	"net/http"
)

func Qr(w http.ResponseWriter, r *http.Request) {
	str.GenerateQRCodeHandler(w, r)
}
