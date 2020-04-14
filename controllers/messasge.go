package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/honkytonktown/GoProject/models"
)

type messageController struct {
	messageIDPattern *regexp.Regexp
}

func (mc messageController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/messages" {
		switch r.Method {
		case http.MethodGet:
			mc.getAll(w, r)
		case http.MethodPost:
			mc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := mc.messageIDPattern.FindStringSubmatch(r.URL.Path)
		fmt.Printf("URL path is: %v\n", r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			mc.get(id, w)
		case http.MethodPut:
			mc.put(id, w, r)
		case http.MethodDelete:
			mc.delete(id, w)
		default:
			w.WriteHeader(http.StatusAccepted)
		}
	}
}

func (mc *messageController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetMessages(), w)
}

func (mc *messageController) get(id int, w http.ResponseWriter) {
	m, err := models.GetMessageByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) post(w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(*r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse object"))
		return
	}
	m, err = models.AddMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) put(id int, w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(*r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not parse object"))
		return
	}
	if id != m.ID {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ID of submitted user must match ID in url"))
		return
	}

	m, err = models.UpdateMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveMessageByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
}

func (mc *messageController) parseRequest(r http.Request) (models.Message, error) {
	dec := json.NewDecoder(r.Body)
	var m models.Message
	err := dec.Decode(&m)
	if err != nil {
		return models.Message{}, err
	}
	return m, nil
}
func newMessageController() *messageController {
	return &messageController{
		messageIDPattern: regexp.MustCompile(`^/messages/(\d+)/?`),
	}
}
