/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
*/
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"sync"
	"time"

	"github.com/shivkumarsingh7/spicework-assignment/handlers"
	"github.com/shivkumarsingh7/spicework-assignment/middleware"
	"github.com/shivkumarsingh7/spicework-assignment/repository"
	"github.com/shivkumarsingh7/spicework-assignment/util"
)

// declaring global variable for thread safe
var lock sync.Mutex

func main() {

	// Initializing HealthCheckRepository
	healthCheckRepo := repository.GetHealthCheckRepository()

	// creating ticker for checking resource health on every 20 seconds
	ticker := time.NewTicker(20000 * time.Millisecond)

	go func() {
		for range ticker.C {
			// locking current operation for thread safe
			lock.Lock()

			// fetcing all resources need to monitor
			for k, v := range util.Resources {
				// converting value to map
				data := v.(map[string]interface{})

				// Check for httpstatus code of given resources
				status := util.CheckResourceAvailabilty(data["url"].(string))

				// setting the new staus of resources
				healthCheckRepo.SetResource(k, status)
			}
			// unlocking current operation
			lock.Unlock()
		}
	}()

	finalHandler := http.HandlerFunc(handlers.Welcome)
	helloworldHandler := http.HandlerFunc(handlers.HelloWorld)

	// initiating routes for root
	http.Handle("/", middleware.Middleware((finalHandler)))

	// initiating routes for hello
	http.Handle("/hello", middleware.Middleware((helloworldHandler)))

	fmt.Println("\nServer stared on port 8080")

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("README")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("Open URL in browser http://localhost:8080")
	fmt.Println("Wait for 20 seconds :-  resources health status will updated")
	fmt.Println("By default all resources health will return 200 status")
	fmt.Println("for testing 500 status please change URL for any resource with wrong path")
	fmt.Println("Edit any resource status/url in spicework-assignment/util/monitor.go")
	fmt.Println(strings.Repeat("=", 80))
	// handling incoming request on given sport
	http.ListenAndServe(":8080", nil)
}
