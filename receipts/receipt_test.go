package receipts

import (
	"testing"
)

// func Test_processReceipt(t *testing.T) {
// 	type args struct {
// 		w http.ResponseWriter
// 		r *http.Request
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			processReceipt(tt.args.w, tt.args.r)
// 		})
// 	}
// }

// func Test_calculatePoints(t *testing.T) {
// 	type args struct {
// 		receipt Receipt
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := calculatePoints(tt.args.receipt); got != tt.want {
// 				t.Errorf("calculatePoints() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
