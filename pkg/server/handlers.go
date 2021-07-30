package server

import (
	"encoding/json"
	"net/http"

	"github.com/erickir/tinyurl/pkg/base62"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Post("/", convertToShortUrl())
	r.Get("/{shortUrl}", getLongUrl())
	return r
}

// Make the api use the mongo db database
type LongUrlRequest struct {
	LongUrl string `json:"longUrl"`
}

type ShortUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}

var urlList []ShortUrlResponse = []ShortUrlResponse{}

func getLongUrl() func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		shortUrl := chi.URLParam(r, "shortUrl")
		var shortUrlResponse ShortUrlResponse
		for _, u := range urlList {
			if u.ShortUrl == shortUrl {
				shortUrlResponse = u
			}
		}

		http.Redirect(rw, r, shortUrlResponse.LongUrl, http.StatusTemporaryRedirect)
	}
}

func convertToShortUrl() func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var longUrl LongUrlRequest

		err := json.NewDecoder(r.Body).Decode(&longUrl)

		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		nextId := int64(hash(longUrl.LongUrl))
		if nextId < 0 {
			nextId *= -1
		}
		shortUrl := base62.ToBase62(nextId)

		var shortUrlResponse ShortUrlResponse
		shortUrlResponse.ShortUrl = shortUrl
		shortUrlResponse.LongUrl = longUrl.LongUrl
		urlList = append(urlList, shortUrlResponse)
		json.NewEncoder(rw).Encode(&shortUrlResponse)
	}
}

func hash(s string) int64 {
	var h int64
	for i := 0; i < len(s); i++ {
		h = h*131 + int64(s[i])
	}
	return h
}
