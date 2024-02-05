package web

import (
	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {
	r := mux.NewRouter()


	r.HandleFunc("/", makeHTTPHandleFunc(homeHandler))
	r.HandleFunc("/login", makeHTTPHandleFunc(loginAPIHandlerGET)).Methods("GET")
	r.HandleFunc("/login", makeHTTPHandleFunc(loginAPIHandlerPOST)).Methods("POST")
	r.HandleFunc("/logout",makeHTTPHandleFunc(logoutHandlerPOST)).Methods("POST")
	r.HandleFunc("/forget",makeHTTPHandleFunc(forgetPasswordHandlerPOST)).Methods("POST")


	return r
}


