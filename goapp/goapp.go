package main

import (
	"fmt"
	"net/http"
	"os"
)

func getMTLS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API\n"))
}

func main() {
	http.HandleFunc("/teste/v1/api/mtls", getMTLS)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
