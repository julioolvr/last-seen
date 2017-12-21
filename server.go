package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	godotenv "gopkg.in/joho/godotenv.v1"
)

type apiResponse struct {
	Response apiData
}

type apiData struct {
	Checkins checkinData
}

type checkinData struct {
	Items []checkinItem
}

type checkinItem struct {
	Venue struct {
		Location struct {
			City    string
			Country string
		}
	}
}

func createHandler(apiURL string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(apiURL)

		if err != nil {
			// TODO: Non-fatal errors
			log.Fatal("Error making request to Foursquare")
		}

		var response apiResponse
		err = json.NewDecoder(resp.Body).Decode(&response)

		if err != nil {
			log.Fatalf("Error decoding JSON %s", err)
		}

		// TODO: Validate presence of location, city, country, etc
		checkin := response.Response.Checkins.Items[0]

		fmt.Fprintf(w, "%s, %s", checkin.Venue.Location.City, checkin.Venue.Location.Country)
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}

	apiURL := fmt.Sprintf("https://api.foursquare.com/v2/users/self/checkins?oauth_token=%s&v=20171220&limit=1&sort=newestfirst", os.Getenv("OAUTH_TOKEN"))

	http.HandleFunc("/", createHandler(apiURL))
	http.ListenAndServe(":8080", nil)
}
