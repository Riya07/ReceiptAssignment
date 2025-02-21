package receipts

type Receipt struct {
	Retailer     string  `json:"retailer"`
	Total        float64 `json:"total"`
	Items        []Item  `json:"items"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price"`
}

type ReceiptResponse struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int `json:"points"`
}
