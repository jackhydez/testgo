package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    w.Write([]byte("hello world 333"))
}

func main() {
    router := httprouter.New()
    router.GET("/", IndexHandler)
    log.Fatal(http.ListenAndServe(":9091", router))
}



