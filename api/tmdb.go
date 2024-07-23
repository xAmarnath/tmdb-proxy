package handler

import (
	"io"
	"net/http"
	"strings"
)

func ImdbProxy(w http.ResponseWriter, r *http.Request) {
	raw_url := r.URL.String()
	raw_url_ := strings.SplitN(raw_url, "/api/tmdb/", 2)
	if len(raw_url_) < 2 {
		raw_url = ""
	} else {
		raw_url = raw_url_[1]
	}

	req, _ := http.NewRequest("GET", "https://api.themoviedb.org/"+raw_url, nil)

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}

	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	return
}
