package handler

import (
	"github.com/AshokShau/upi-qr-generator/str"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	str.Home(w, r)
}
