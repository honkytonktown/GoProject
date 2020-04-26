package controllers

import (
	"fmt"
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
	case http.MethodPost:
		pc.Post(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (pc *postController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetPosts(), w)
}

func (pc *postController) Post(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetPosts(), w)
}
func newPostController() *postController {
	fmt.Println("postController created")
	//establish connection to db
	return &postController{
		messageIDPattern: regexp.MustCompile(`^/posts/(\d+)/?`),
	}
}
