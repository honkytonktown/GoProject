package controllers

import (
	"net/http"
	"regexp"

	"github.com/honkytonktown/GoProject/models"
)

type postController struct {
	messageIDPattern *regexp.Regexp
}

func (pc postController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	models.Connect()
	switch r.Method {
	case http.MethodGet:
		pc.getAll(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (pc *postController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetPosts(), w)
}
func newPostController() *postController {
	return &postController{
		messageIDPattern: regexp.MustCompile(`^/posts/(\d+)/?`),
	}
}
