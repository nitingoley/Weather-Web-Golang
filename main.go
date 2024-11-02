package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Struct to hold the API configuration data
type apiConfigData struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// Struct to hold weather data from the API response
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

// Function to load API config from a specified JSON file
func loadApiConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return apiConfigData{}, err
	}

	var config apiConfigData
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return apiConfigData{}, err
	}

	log.Printf("API Key loaded successfully.")
	return config, nil
}

// Function to query the weather data for a given city
func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, fmt.Errorf("could not load API config: %v", err)
	}

	// Build the URL for the weather API request
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?APPID=%s&q=%s", apiConfig.OpenWeatherMapApiKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return weatherData{}, fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response status
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf("API request failed. Status: %d. Body: %s", resp.StatusCode, body)
		return weatherData{}, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Decode JSON response into weatherData struct
	var data weatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return weatherData{}, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	log.Printf("Weather data for %s: %+v", city, data) // Debug log of weather data
	return data, nil
}

// Simple hello handler for testing server
func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go!\n"))
}

// Weather handler to fetch weather data for a specified city
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2] // Extract city from URL

	data, err := query(city)
	if err != nil {
		log.Printf("Error querying data for city %s: %v", city, err)
		http.Error(w, "Could not fetch weather data", http.StatusInternalServerError)
		return
	}

	// Set content type to JSON and encode the weather data to response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}

// serve html file in golang project

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/weather/", weatherHandler)
	http.HandleFunc("/", indexHandler)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
