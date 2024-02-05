package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type handlers func(w http.ResponseWriter, r *http.Request) error

func makeHTTPHandleFunc(h handlers) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w,r); err != nil {
			log.Println(err)
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	
	switch v := v.(type){
	case ApiError:
		json.NewEncoder(w).Encode(ApiError{ApiError: v.String()})
		log.Println(v.String())
	case string:
		json.NewEncoder(w).Encode(ApiError{ApiError: v})
		log.Println(v)
		
	}
	return nil
}

func (a ApiError) String() string {
	return a.ApiError
}

type ApiError struct {
	ApiError string `json:"ApiError"`
}