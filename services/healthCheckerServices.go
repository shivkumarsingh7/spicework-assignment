/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- services
| description:-  service contains of business logic
*/
package services

import (
	"github.com/shivkumarsingh7/spicework-assignment/repository"
)

// GetHealthResource : check for all monitored resources are healthy or not
func GetHealthResource() bool {
	healthCheckRepo := repository.GetHealthCheckRepository()
	status := healthCheckRepo.IsHealthy()
	return status
}
