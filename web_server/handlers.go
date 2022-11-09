package main

import (
	"fmt"
	"net/http"
)

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello word")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the API EndPoint")
}
