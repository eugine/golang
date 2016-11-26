package main

import (
	"fmt"
	"log"
	"net/http"
)

type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}
