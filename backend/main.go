package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port = os.Getenv("FRONTEND_PORT")
)

func init() {
	flag.StringVar(&port, "port", port, "Port to run on")
	flag.Parse()
}

type User struct {
	ID       uint
	PreName  string
	LastName string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Start work..."))
		time.Sleep(10 * time.Second)
		w.Write([]byte("Finished work!"))
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(User{1, "John", "Doe"}); err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	})

	listenAddr := fmt.Sprintf(":%v", port)
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}
