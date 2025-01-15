package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func handleWork(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Start work..."))
	time.Sleep(10 * time.Second) // simulate work
	w.Write([]byte("Finished work!"))
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	// get user id from url params
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// find user by id
	user := FindUserByID(uint(id))
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// respond with user json
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func handleGetUserList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
