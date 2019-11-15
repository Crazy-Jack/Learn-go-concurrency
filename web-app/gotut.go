package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) { // use *http.Request to read through the large file without load it.
	fmt.Fprint(w, "WOW, Go is super interesting!")

}

func main() {
	// route the url
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8000", nil)
}
