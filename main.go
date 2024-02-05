package main

import (
	"github.com/projectX/service/web"
)

func main() {
	router := web.NewRouter()
	web.InitServer(":4000", router)
}

