package main

import (
	router "gihub.com/imyashkale/go-basic/Router"
	// Mux "gihub.com/imyashkale/go-basic/ServeMux"
)

func main() {

	// ServeMux Example
	// Mux.StartServrMux()
	// Start Custome Mx
	// Mux.StartCustomMx()
	// Multiple handlers => Multiple Routes
	//Mux.MultipleHandlers()
	// router.HTTPRouter()
	router.FileServer()
}
