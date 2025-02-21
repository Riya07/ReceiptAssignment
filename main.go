package main

import (
	"ReceiptAssignment/receipts"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/receipts/process", receipts.ProcessReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", receipts.GetPoints).Methods("GET")

	log.Println("Listening at :8080")
	http.ListenAndServe(":8080", r)

}
