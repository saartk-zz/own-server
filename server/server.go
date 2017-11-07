package main

import (
	"fmt"
	"net/http"
)

func HandleFunction(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Yosi")
}

func main() {
	http.HandleFunc("/",HandleFunction)
	http.ListenAndServe(":50001",nil)
}