package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Sample Golang app running...")
	message := os.Getenv("MSG_ENV")
	if message == "" {
		message = ":( MSG_ENV variable not defined"
	}
	// added 20 second delay
	time.Sleep(20 * time.Second)
	fmt.Fprintf(w, "<h1>%s</h1>", message)
}

func main() {
	flag.Parse()
	log.Print("Sample Golang app server started...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
