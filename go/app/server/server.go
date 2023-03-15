package server

import (
	"encoding/json"
	"net/http"
)

type Ping struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

func TestGet(w http.ResponseWriter, r *http.Request) {
	tPing := &Ping{http.StatusOK, "ok"}

	tPingByte, tError := json.Marshal(tPing)
	if tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	if _, tError := w.Write(tPingByte); tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func TestPost(w http.ResponseWriter, r *http.Request) {
	tBody := make([]byte, r.ContentLength)
	r.Body.Read(tBody)

	tPing := &Ping{}
	if tError := json.Unmarshal(tBody, tPing); tError != nil {
		http.Error(w, tError.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SignUp(w http.ResponseWriter, r *http.Request) {

}
