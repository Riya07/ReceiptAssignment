package receipts

import (
	"testing"
)

// Test Points Calculation Logic
func TestCalculatePoints(t *testing.T) {
	testReceipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            6.49,
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            12.25,
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            1.26,
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            3.35,
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            12.00,
			},
		},
		Total: 35.35,
	}

	calculatedPoints := calculatePoints(testReceipt)
	expectedPoints := 28
	if calculatedPoints != expectedPoints {
		t.Fatalf("Expected %d points, but got %d", expectedPoints, calculatedPoints)
	}
}
