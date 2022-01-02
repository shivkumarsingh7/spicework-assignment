/*
|-----------------------------------------------------------------------
| Spicework Assignment
| author:- shiv kumar singh <shivam.developer7@gmail.com>
| package:- repository
| description:-  repository has business logic & database operation
*/
package repository

import (
	"fmt"
	"strings"

	"github.com/shivkumarsingh7/spicework-assignment/util"
)

// HealthAggregator : interface with set of methods to be implemented
type HealthAggregator interface {
	SetResource(resourceName string, isHealthy bool)
	IsHealthy() bool
}

type HealthCheckRepository struct{}

// global declration of health repository
var healthCheckRepo *HealthCheckRepository

// GetHealthCheckRepository : return healthCheckRepository instance
func GetHealthCheckRepository() HealthAggregator {
	if healthCheckRepo == nil {
		healthCheckRepo = &HealthCheckRepository{}
	}
	return healthCheckRepo
}

// IsHealthy :  Check for health of all avialable resources and return boolean
func (healthCheckRepo *HealthCheckRepository) IsHealthy() bool {
	// formating output on terminal
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("RESOURCE HEALTH")
	fmt.Println(strings.Repeat("=", 80))

	// traverse through each resources
	for _, v := range util.Resources {
		data := v.(map[string]interface{})
		status := data["status"].(bool)

		// If any one of the resources is unhealthy, return false
		if !status {
			fmt.Println(data["title"], " : Unhealthy")
			return status
		}
		fmt.Println(data["title"], " : Healthy")
	}
	return true
}

// SetResource : setting staus for resources
func (healthCheckRepo *HealthCheckRepository) SetResource(resourceName string, isHealthy bool) {
	fmt.Println("Set Resource ", resourceName, " Status ", isHealthy)
	data := util.Resources[resourceName].(map[string]interface{})
	data["status"] = isHealthy
}
