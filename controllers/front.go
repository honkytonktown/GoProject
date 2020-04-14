package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	mc := newMessageController()
	pc := newPostController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	http.Handle("/messages", *mc)
	http.Handle("/messages/", *mc)

	http.Handle("/posts", *pc)
	http.Handle("/posts/", *pc)

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
