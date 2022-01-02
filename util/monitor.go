/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- util
| description:- util library is used to define resources to be monitored
*/
package util

import (
	"log"
	"net/http"
)

// Resources : defining global map
var Resources map[string]interface{}

// initializing resources to be monitored
func init() {

	// given resources are health monitored
	// we can monitore any 3rd party resource like css,js,font,database etc
	// but for demo purpose we are using css,js and URLs and check for httpstatus
	Resources = map[string]interface{}{
		"bs": map[string]interface{}{
			"title":  "BootStrap",
			"url":    "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js",
			"status": true,
		},
		"bulma": map[string]interface{}{
			"title":  "Bulma",
			"url":    "https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.5/css/bulma.min.css",
			"status": true,
		},
		"golang": map[string]interface{}{
			"title":  "Golang Doc",
			"url":    "https://golang.org/doc/",
			"status": true,
		},
		"youtube": map[string]interface{}{
			"title":  "Youtube",
			"url":    "https://www.youtube.com/watch?v=s_gRXOsrDAA",
			"status": true,
		},
	}
}

// CheckResourceAvailabilty : checking for httpstatus code 200
// and returning true else false
func CheckResourceAvailabilty(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return true
	}
	return false
}
