package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func dec_to_base62(symbol string, n int) string {
	var url string
	for int(n/62) > 0 {
		url += string(symbol[n%62])
		n = int(n / 62)
	}
	url += string(symbol[n])
	return reverse(url)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://google.com", http.StatusMovedPermanently)
	fmt.Println(mux.Vars(r)["id"])
}
func create(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create")
}
func info(w http.ResponseWriter, r *http.Request) {

	fmt.Println("info")
}
func view(w http.ResponseWriter, r *http.Request) {

	fmt.Println("view")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", view).Methods("GET")
	r.HandleFunc("/", create).Methods("POST")
	r.HandleFunc("/{id:[a-zA-Z0-9]+}", info).Methods("POST")
	r.HandleFunc("/{id:[a-zA-Z0-9]+}", redirect).Methods("GET")

	fmt.Println(dec_to_base62("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 3843))

	http.Handle("/", r)
	http.ListenAndServe(":"+"8000", nil)

}
