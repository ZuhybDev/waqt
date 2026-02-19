package handlers

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func HandlePrayers(w http.ResponseWriter, r *http.Request) {

	// base url
	baseURL := "https://api.aladhan.com/v1/timingsByCity/"

	//validate the meth
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}

	//query params to find specific country and city
	city := r.URL.Query().Get("city")
	country := r.URL.Query().Get("country")

	fmt.Printf("city %s country %s\n", city, country)

	// validate queries
	if city == "" || country == "" {
		http.Error(w, "city and country are required", http.StatusBadRequest)
		return
	}

	date := time.Now().Format("02-01-2006")
	// merge the URL
	fullURL := fmt.Sprintf("%s%s?city=%s&country=%s", baseURL, date, city, country)

	fmt.Printf("full url: %s\n", fullURL)

	// new http client for request
	client := &http.Client{}

	// new request init
	req, err := http.NewRequest("GET", fullURL, nil)

	// handle error gracefully
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	/// sending request
	resp, err := client.Do(req)

	// handle error gracefully
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	// HTTP headers
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	// finally return the response
	w.Write(body)

}

func HandleQuran(w http.ResponseWriter, r *http.Request) {


  
}






