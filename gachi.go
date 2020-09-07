package main

import (
	"fmt"
	"net/http"
)

func getGachiPhrase(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Fucking cumming")
}

func main() {
	http.HandleFunc("/", getGachiPhrase)
    http.ListenAndServe(":8080", nil)

}