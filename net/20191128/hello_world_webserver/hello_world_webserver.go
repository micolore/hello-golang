package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {

	fmt.Println("inside helloserver handler")
	fmt.Fprintf(w, "hello,"+req.URL.Path[1:])
}

func main() {

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("listenAndServe:", err.Error())
	}

}