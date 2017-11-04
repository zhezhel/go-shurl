package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"time"

	"./model"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const urlSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var intMap map[int]string
var strMap map[string]int

func reverse(value string) string {
	data := []rune(value)
	result := []rune{}

	for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
	}
	return string(result)
}

func shuffle(s string) (map[int]string, map[string]int) {
	rand.Seed(time.Now().UnixNano())

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

func getUrlFromDb(db *gorm.DB, shortUrl string) (string, error) {
	fmt.Println(shortUrl)
	url := model.Url{}
	err := db.Where("short_url = ?", shortUrl).First(&url).Error
	if err == gorm.ErrRecordNotFound {
		return "", err
	}
	return url.LongUrl, nil
}

func createUrlInDb(db *gorm.DB, LongUrl string, id int) (string, error) {

	url := model.Url{
		ShortUrl:     decToBase62(intMap, id),
		LongUrl:      string(LongUrl),
		Redirections: 0,
	}

	err := db.Save(&url).Error
	if err != nil {
		return "", err
	}
	return url.ShortUrl, nil

}

func redirect(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := mux.Vars(r)["shortUrl"]
		longUrl, err := getUrlFromDb(db, shortUrl)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}
		http.Redirect(w, r, longUrl, http.StatusMovedPermanently)
	}
}

func create(db *gorm.DB, id int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, _ := ioutil.ReadAll(r.Body)
		shortUrl, err := createUrlInDb(db, string(data), id)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id = id + 1
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"short_url\": \"" + shortUrl + "\" }"))
	}
}

// TODO: Rewrite with using json
func info(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("info")
		vars := mux.Vars(r)
		shortUrl := vars["shortUrl"]
		longUrl, err := getUrlFromDb(db, shortUrl)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(longUrl))

	}
}

func view(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view")
}

func getRouter(db *gorm.DB) *mux.Router {
	// Init
	id := 0
	intMap, strMap = shuffle(urlSymbols)
	url := model.Url{}
	err := db.Last(&url).Error
	if err != gorm.ErrRecordNotFound {
		id = base62ToDec(strMap, url.ShortUrl) + 1
	}

	r := mux.NewRouter()
	r.HandleFunc("/", view).Methods("GET")
	r.HandleFunc("/", create(db, id)).Methods("POST")
	r.HandleFunc("/{shortUrl:[a-zA-Z0-9]+}", info(db)).Methods("POST")
	r.HandleFunc("/{shortUrl:[a-zA-Z0-9]+}", redirect(db)).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}
