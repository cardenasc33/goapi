package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi" //package for web dev
	"github.com/cardenasc33/goapi/internal/handlers" // importing our own package
	log "github.com/sirupsen/logrus" // used to log errors for debugging

)

func main () {
	log.SetReportCaller(true) // When printing something out, you get the file and line number
	
	// Returns a pointer to a MUX type
	var r *chi.Mux = chi.NewRouter()  //struct to set up API
	handlers.Handler(r)  // Set up router, i.e add enpoint definitions

	fmt.Println("Starting Go API service...")

	// Start server
	// @params: (base location of server, handler that Mux type satisfies)
	var ip = "localhost"
	var port = "8000"
	var socket = ip + ":" + port 
	fmt.Println("Serving on ", socket)
	err := http.ListenAndServe( socket , r)
	if err != nil {
		log.Error(err)  // log any errors when starting the server
	}
	
	
	
}