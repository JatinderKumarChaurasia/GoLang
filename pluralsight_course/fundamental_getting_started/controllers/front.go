package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	userCtrl := newUserController()
	http.Handle("/users", *userCtrl)
	http.Handle("/users/", *userCtrl)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	_ = enc.Encode(data)
}
