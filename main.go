package main

import (
	"log"
	"net/http"
    "fmt"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    w.Write([]byte("hello world 777"))
}

func main() {
    fmt.Println("hello 1")
    router := httprouter.New()
    router.GET("/", IndexHandler)
    log.Fatal(http.ListenAndServe(":9091", router))
}



