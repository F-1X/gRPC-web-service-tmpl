package web

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// init of local memory pool of users
// just want to get it from frontend from forms register/authentification
var memoryCreds = make(map[string]string)

