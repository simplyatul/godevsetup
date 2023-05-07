package main

import (
	"fmt"
	"os"
)

func getHostName() string {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hn
}

func main() {
	hn := getHostName()
	fmt.Printf("Hostname is : %s\n", hn)
}
