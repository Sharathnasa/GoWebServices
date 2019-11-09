package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterController() {

	uc := newUserController()

	http.Handle("/user", uc)
	http.Handle("/user/", uc)
}

func encodeResponseAsJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
