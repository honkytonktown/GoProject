package controllers

import (
	"net/http"
	"regexp"
)

type SQLController struct {
	messageIDPattern *regexp.Regexp
}

func (sd SQLController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
func newSQLController() *SQLController {
	return &SQLController{
		messageIDPattern: regexp.MustCompile(`^/sqldata/(\d+)/?`),
	}
}
