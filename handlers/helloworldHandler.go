/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- handlers
| description:- endpoint for incoming requests
*/
package handlers

import "net/http"

// HelloWorld :  returing Json response to application with status code 200
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	ToJSON(w, http.StatusOK, struct{ Message string }{Message: "Hello World!!"})
}
