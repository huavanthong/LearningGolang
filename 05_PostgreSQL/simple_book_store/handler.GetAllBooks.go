package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
