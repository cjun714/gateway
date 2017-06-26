package main

import (
	"net/http"
	"os"
	"util/log"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/token", token)
	router.GET("/validate", validate)

	log.H("Server is started @:", os.Args[1])
	log.E(http.ListenAndServe(":"+os.Args[1], router))
}
