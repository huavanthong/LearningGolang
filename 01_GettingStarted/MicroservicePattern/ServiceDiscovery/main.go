package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Service struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Services struct {
	sync.Mutex
	data []*Service
}

func main() {
	var services Services

	// Register a new service
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var service Service

		if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		services.Lock()
		defer services.Unlock()

		// Check if the service is already registered
		for _, s := range services.data {
			if s.ID == service.ID {
				w.WriteHeader(http.StatusConflict)
				fmt.Fprintf(w, "service with ID %s already exists", service.ID)
				return
			}
		}

		// Add the service to the list
		services.data = append(services.data, &service)

		fmt.Fprintf(w, "registered service %s at %s:%d", service.Name, service.Address, service.Port)
	})

	// List all registered services
	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		services.Lock()
		defer services.Unlock()

		json.NewEncoder(w).Encode(services.data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
