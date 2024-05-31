package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type URL struct {
	Original string `json:"original"`
	Short    string `json:"short"`
}

var (
	urls    = make(map[string]string)
	urlMux  sync.Mutex
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rng     = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL()
	urlMux.Lock()
	urls[shortURL] = requestData.URL
	urlMux.Unlock()

	responseData := URL{
		Original: requestData.URL,
		Short:    shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]

	urlMux.Lock()
	originalURL, exists := urls[shortURL]
	urlMux.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

func generateShortURL() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rng.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURLHandler).Methods("POST")
	r.HandleFunc("/{shortURL}", redirectHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
