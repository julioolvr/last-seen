package oauth_server

import (
	"fmt"
	"net/http"
)

// Quick server to get the token from Foursquare's Oauth flow
// Flow instructions at
// https://developer.foursquare.com/docs/api/configuration/authentication

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got code: %s", r.URL.Query()["code"][0])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
