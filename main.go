package main

import (
	"log"
	"net/http"
    "fmt"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
    w.Write([]byte("hello world 556"))
}

func main() {
    fmt.Println("hello")
    router := httprouter.New()
    router.GET("/", IndexHandler)
    log.Fatal(http.ListenAndServe(":9091", router))
}



