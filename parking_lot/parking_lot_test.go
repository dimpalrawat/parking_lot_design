package parking_lot

import "testing"

func TestMainParkingLot(t *testing.T) {
	result := MainParkingLot("C")
	if result != false {
		t.Errorf("Error TestMainParkingLot Expected: %v Actual: %v", false, result)
	}
}