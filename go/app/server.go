package main

import (
	"net/http"
)

func TestServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
