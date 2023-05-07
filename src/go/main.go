package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

func getHostName() string {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname is : %s\n", hn)
	return hn
}

func handleRoot(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hostname: %s\n", getHostName())
}

func handleRequests() {
    http.HandleFunc("/", handleRoot)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
