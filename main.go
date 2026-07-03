package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to my golang server")
}

func submit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "this method is not valid", http.StatusMethodNotAllowed)
	}
	return
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/submit", submit)

	fmt.Println("server running on port :9090")
	http.ListenAndServe(":9090", nil)
}
