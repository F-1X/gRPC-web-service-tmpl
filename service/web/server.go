package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func InitServer(l string, router *mux.Router) {
	log.Fatal(http.ListenAndServe(l, router))
}