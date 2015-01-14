package main

import (
	"github.com/mathcunha/bandeiraenergia/band"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", band.DoRequest)
	probe := band.Probe{}
	probe.Run()
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: 0.0.0.0:8080 ", err)
	}
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
