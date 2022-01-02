/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- handlers
| description:- endpoint for incoming requests
*/
package handlers

import (
	"net/http"
)

// GetResource :  returing Json response to application with status code 200
func Welcome(w http.ResponseWriter, r *http.Request) {
	ToJSON(w, http.StatusOK, struct{ Message string }{Message: "Welcome to the application !!"})
}
