/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- middleware
| description:-  middleware used for performing set of task before
| passing data to application.
*/
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/shivkumarsingh7/spicework-assignment/services"
)

// Middleware : method used for checking resource health whenever
// for new incomings request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check for all resources health
		status := services.GetHealthResource()

		// If any one of the resources is unhealthy, your web application
		// will respond to the web request with a 5XX code.
		if !status {
			w.Header().Add("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusServiceUnavailable)
			data := struct{ Message string }{Message: "503 Service Unavailable"}
			json.NewEncoder(w).Encode(data)
			return
		}
		// else normal execution of application
		next.ServeHTTP(w, r)
	})
}
