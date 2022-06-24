package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	kvs "github.com/nathanielfernandes/kv/lib/kvserver"
)

func main() {
	kvs := kvs.NewKVServer()

	router := httprouter.New()
	router.GET("/kv/:key", kvs.Get)
	router.POST("/kv/:key/:value", kvs.Set)

	router.GET("/r/:key", kvs.RedirectTo)

	fmt.Printf("KV\nListening on port 80\n")
	if err := http.ListenAndServe("0.0.0.0:80", router); err != nil {
		log.Fatal(err)
	}
}
