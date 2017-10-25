package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

func getHandler(status int, message string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(message))
	}
}

func main() {
	port := flag.String("port", "8080", "the port")
	message := flag.String("message", "500 - not accessible over a public connection", "message to return")
	statusCode := flag.Int("status", 500, "status code to return")
	log.Println("parsing arguments...")
	flag.Parse()
	log.Println("port", *port)
	log.Println("message", *message)
	log.Println("statusCode", *statusCode)
	log.Println("creating server...")
	s := &http.Server{
		Addr:           ":" + *port,
		Handler:        http.HandlerFunc(getHandler(*statusCode, *message)),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("serving on localhost" + s.Addr)
	log.Fatal(s.ListenAndServe())
}
