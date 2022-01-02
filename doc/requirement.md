# Spiceworks Inc
## P&E HOMEWORK - Health Checker
HI! Please review the homework question. Complete the homework and save the file to your home drive. At the bottom of this questionnaire, select "choose file" to upload your completed homework. If you have multiple files to upload please do so by answering the additional questions. You may need to zip or tar your files first, and then upload.

You own a web application that receives thousands of requests per second.

* Each web request is evaluated concurrently.

* The application depends on several external resources.

* Each of those resources might be in a healthy or unhealthy state.

* Each resource will be monitored (polled) at different intervals in independent threads.

* Assume the health of each resource will be random and change several times throughout the life of the application.

* The health of each resource will be reported to a central `HealthAggregator` using the method `SetResource`—described below—and will contain the most recent health status for a single resource.

* On each request to the main app, one of the steps to perform is call `HealthAggregator.IsHealthy()` to determine the overall health of the application.

* If any one of the resources is unhealthy, your web application will respond to the web request with a 5XX code.

* Only if all resources are healthy will the response code be a 2XX.

Write a library with a class or struct that fulfills the following interface, using only native (standard) libraries. You may use any language you wish, but it must support thread safety without using the `atomic` package.

Class HealthAggregator

// SetResource will be called at regular intervals by external threads to report one resource’s health

SetResource(resourceName string, isHealthy bool) => returns nothing

// `IsHealthy()` must only be true when all monitored resources are healthy.

// `IsHealthy()` must return false if one or more resources are unhealthy.

IsHealthy() => returns bool
