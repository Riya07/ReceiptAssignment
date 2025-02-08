package main

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	receipts = make(map[int]Receipt)
	points   = make(map[int]int)
	userID   = 0
	mutex    sync.Mutex
)

// ProcessReceipt
func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id := userID
	pointValue := calculatePoints(receipt)

	mutex.Lock()
	receipts[id] = receipt
	points[id] = pointValue
	mutex.Unlock()

	response, _ := json.Marshal(ReceiptResponse{ID: id})
	w.Write(response)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Wrong ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	pointValue, exists := points[userID]
	mutex.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(PointsResponse{Points: pointValue})
	w.Write(response)
}

func calculatePoints(receipt Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			points++
		}
	}

	// Rule 2: 50 points if total is a round dollar amount
	if receipt.Total == math.Floor(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points per item in receipt
	points += len(receipt.Items) * 5

	// Rule 5: If item description length is a multiple of 3, award price * 0.2 points
	for _, item := range receipt.Items {
		if len(item.ShortDescription)%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	return points
}
