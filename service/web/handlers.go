package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)


func homeHandler(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(sessionKey)
	if err != nil {
		http.Redirect(w,r, "/login",http.StatusUnauthorized)
		return nil

	}

	session, ok := sessions[cookie.Value]
	if !ok {
		http.Redirect(w,r,  "/login",http.StatusSeeOther)
		return nil
	
	}

	if time.Now().After(session.Expires) {
		log.Println("exist but expired")
		http.Redirect(w,r, "/login",http.StatusSeeOther)
		return nil
		
	}

	return WriteJSON(w, http.StatusAccepted, "session success")
}


func init() {
	memoryCreds["gleb"] = "password"
}

func loginAPIHandlerGET(w http.ResponseWriter, r *http.Request) error {
	data, err := os.ReadFile("index.html")
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, err)
	}

	w.Write(data)
	return nil
}

func loginAPIHandlerPOST(w http.ResponseWriter, r *http.Request) error {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return WriteJSON(w, http.StatusBadRequest, ApiError{"err decode creds"})
	}
	password, ok := memoryCreds[creds.Username]
	if !ok {
		return WriteJSON(w, http.StatusBadRequest, ApiError{"not exist user"})
	}
	if password != creds.Password {
		return WriteJSON(w, http.StatusUnauthorized, ApiError{"wrong password"})
	}
	sessindId := uuid.NewString()
	expires := time.Now().Add(5 * time.Minute) // 5 minutes
	maxAge := 1 * 60 * 60 // 1 hour
	sessions[sessindId] = sessionCookie{
		Name:	sessionKey,
		Value:   sessindId,
		Expires: expires,
		MaxAge:  maxAge,
	}
	cookie := &http.Cookie{
		Name: sessionKey,
		Value: sessindId, 
		Expires: expires,
		MaxAge:  maxAge,
	}
	http.SetCookie(w, cookie)
	return WriteJSON(w, http.StatusOK, "success authorizated")
}


func forgetPasswordHandlerPOST(w http.ResponseWriter, r *http.Request) error {
	fmt.Println(r.Body)
	return nil
}	

func logoutHandlerPOST(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(sessionKey)
	if err != nil {
		return nil
	}

	delete(sessions, cookie.Value)
	http.SetCookie(w, &http.Cookie{
		Name:   sessionKey,
		MaxAge: -1,
	})
	
	return WriteJSON(w, http.StatusAccepted, ApiError{"logout success"})
}
