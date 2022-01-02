/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- handlers
| description:- endpoint for incoming requests
*/
package handlers

import (
	"encoding/json"
	"net/http"
)

const jsonContentType = "application/json; charset=utf-8"

// ToJSON : helper function for json response
func ToJSON(w http.ResponseWriter, statusCode int, obj interface{}) error {
	w.Header().Add("Content-Type", jsonContentType)
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(obj)
}
