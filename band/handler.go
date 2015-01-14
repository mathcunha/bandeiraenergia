package band

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var currentStatus []Status

func Update(status []Status) {
	currentStatus = status
}

func DoRequest(w http.ResponseWriter, r *http.Request) {
	if "GET" == r.Method {
		if strings.Contains(r.RequestURI, "pretty") {
			Pretty(w, currentStatus)
		} else {
			Print(w, currentStatus)
		}
	}
}

func Print(w http.ResponseWriter, status []Status) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentStatus)
}

func Pretty(w http.ResponseWriter, status []Status) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", "[")
	length := len(status) - 1
	for i, s := range status {

		fmt.Fprintf(w, "{\"Region\":\"%v\",\"Color\":\"%v\"}", s.Region, s.Color)
		if i != length {
			fmt.Fprintf(w, "%v", ",")
		}
	}
	fmt.Fprintf(w, "%v", "]")
}
