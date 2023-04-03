package handlers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Messages []string `json:"messages"`
	Code     string   `json:"code"`
}

func BadRequest(w http.ResponseWriter, code string, messages ...string) {
	err := &errorResponse{
		Code:     code,
		Messages: messages,
	}
	Render(w, err, http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error) {
	Render(w, &errorResponse{
		Code:     "INTERNAL_SERVER_ERROR",
		Messages: []string{err.Error()},
	}, http.StatusInternalServerError)
}

func Render(w http.ResponseWriter, obj interface{}, status int) {
	js, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func OK(w http.ResponseWriter, obj interface{}) {
	Render(w, obj, http.StatusOK)
}

func NotFound(w http.ResponseWriter, messages ...string) {
	err := &errorResponse{
		Code:     "NOT_FOUND",
		Messages: messages,
	}
	Render(w, err, http.StatusNotFound)
}
