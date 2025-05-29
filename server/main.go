package main

import (
	"fmt"
	"net/http"
)

func myAwesomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, my dudes")
}

func main() {
	http.HandleFunc("/", myAwesomeHandler)	
	http.ListenAndServe(":8080", nil)
}



