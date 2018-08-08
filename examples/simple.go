package main

import (
	"fmt"

	"github.com/nuvi/healthz"
)

func main() {

	healthz.Serve("localhost:8080", "/healthz") // params shown here are redundant, these are the package defaults
	healthError := healthz.HealthError{Description: "Something fatal went wrong!"}
	healthz.NewFatalError(healthError)
	healthError = healthz.HealthError{Description: "This time it wasn't fatal!"}
	healthz.NewNonFatalError(healthError)

	fmt.Println("Serving healthz on http://localhost:8080/healthz")

	forever := make(chan bool)
	<-forever
}
