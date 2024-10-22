package main

import (
	router "goProject/routes"
)

func main() {

	// Create a new Gin router
	router := router.SetupRouter()
	router.Run(":8080")
}
