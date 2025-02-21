package receipts

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	receipts = make(map[string]Receipt)
	points   = make(map[string]int)
	mutex    sync.Mutex
)

// ProcessReceipt
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	receiptId := uuid.New().String()
	pointValue := calculatePoints(receipt)

	mutex.Lock()
	receipts[receiptId] = receipt
	points[receiptId] = pointValue
	mutex.Unlock()

	response, _ := json.Marshal(ReceiptResponse{ID: receiptId})

	w.Write(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	mutex.Lock()
	pointValue, exists := points[receiptID]
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

	// Rule 2: 50 points if total is a round dollar amount with no cents
	if receipt.Total == math.Floor(receipt.Total) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if math.Mod(receipt.Total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points per 2 item in receipt
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: If item description length is a multiple of 3, award price * 0.2 points
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			points += int(math.Ceil(item.Price * 0.2))
		}
	}

	//TODO : How do we ensure it will be a large language model is it a key in the receipt? Who checks it ?
	// Rule 6: 5 points if the total is greater than 10.00
	// if receipt.Total > 10.00 {
	// 	points += 5
	// }

	// Rule 7: 6 points if the day in the purchase date is odd
	dateParts := strings.Split(receipt.PurchaseDate, "-")
	if len(dateParts) == 3 {
		day, err := strconv.Atoi(dateParts[2])
		if err == nil && day%2 == 1 {
			points += 6
		}
	}

	// Rule 8: 10 points if the time of purchase is after 2:00pm and before 4:00pm
	timeParts := strings.Split(receipt.PurchaseTime, ":")
	if len(timeParts) == 2 {
		hour, err := strconv.Atoi(timeParts[0])
		if err == nil && hour >= 14 && hour < 16 {
			points += 10
		}
	}
	return points
}
