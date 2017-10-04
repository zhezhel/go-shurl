package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

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

func shuffle(s string) []string {
	tmp := make([]string, len(s))
	for i := range s {
		tmp[i] = string(s[i])
	}
	for i := range tmp {
		j := rand.Intn(i + 1)
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return tmp
}

func decToBase62(symbol []string, n int) string {
	var url string
	for int(n/62) > 0 {
		url += symbol[n%62]
		n = int(n / 62)
	}
	url += symbol[n]
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

	urlSymbols := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	t := shuffle(urlSymbols)
	fmt.Println(decToBase62(t, 0))

	fmt.Println(t)

	http.Handle("/", r)
	http.ListenAndServe(":"+"8000", nil)

}
