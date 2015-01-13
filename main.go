package main

import (
	"github.com/mathcunha/bandeiraenergia/band"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", band.DoRequest)
	probe := band.Probe{}
	probe.Run()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: 0.0.0.0:8080 ", err)
	}
}
