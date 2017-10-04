package main

import (
	"fmt"
	"math"
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

func shuffle(s string) (map[int]string, map[string]int) {
	tmpIntString := make(map[int]string, len(s))
	tmpStringInt := make(map[string]int, len(s))
	for i := range s {
		tmpIntString[i] = string(s[i])
	}
	for i := range tmpIntString {
		j := rand.Intn(i + 1)
		tmpIntString[i], tmpIntString[j] = tmpIntString[j], tmpIntString[i]
	}
	for k, v := range tmpIntString {
		tmpStringInt[v] = k
	}

	return tmpIntString, tmpStringInt
}

func decToBase62(symbol map[int]string, n int) string {
	var url string
	for int(n/62) > 0 {
		url += symbol[n%62]
		n = int(n / 62)
	}
	url += symbol[n]
	return reverse(url)
}

func base62ToDec(symbol map[string]int, url string) int {
	var num float64
	url = reverse(url)
	for k, v := range url {
		num += float64((symbol[string(v)])) * math.Pow(62, float64(k))
	}
	return int(num)
}

func getUrlFromDb(shortUrl string) string {
	fmt.Println(shortUrl)
	return "https://google.com"
}

func createUrlInDb(longUrl string) string {
	return "s"
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, mux.Vars(r)["id"], http.StatusMovedPermanently)
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
	i, s := shuffle(urlSymbols)

	n := 19

	fmt.Println(decToBase62(i, n))
	tmp := decToBase62(i, n)
	fmt.Println(base62ToDec(s, tmp))

	http.Handle("/", r)
	http.ListenAndServe(":"+"8000", nil)

}
