package parking_lot

import "testing"

func TestParkVehicle(t *testing.T) {
	parkingLotSize := 2
	parkingLot := ParkingLot{
		RegToSlotNoMap:   make(map[string]int),
		BookedSlots:      make([]*Vehicle, parkingLotSize),
		VacatedSlots:     &VacatedSlotsHeap{},
	}
	parkingLot.ParkingSlotSize = parkingLotSize
	parkingLot.VacatedSlots.InitializeHeap(parkingLotSize)
	vehicle1 := Vehicle{
		Color:     "White",
		RegNumber: "KA-01-HH-1234",
	}
	slot1 := parkingLot.ParkVehicle(vehicle1)
	if slot1 != 1 {
		t.Errorf("Error ParkVehicle: Expected: %v Actual: %v", 1, slot1)
	}
	vehicle2 := Vehicle{
		Color:     "Black",
		RegNumber: "KA-01-BB-0001",
	}
	slot2 := parkingLot.ParkVehicle(vehicle2)
	if slot2 != 2 {
		t.Errorf("Error ParkVehicle: Expected: %v Actual: %v", 2, slot2)
	}
	vehicle3 := Vehicle{
		Color:     "red",
		RegNumber: "KA-01-CC-0231",
	}
	slot3 := parkingLot.ParkVehicle(vehicle3)
	if slot3 != -1 {
		t.Errorf("Error ParkVehicle: Expected: %v Actual: %v", -1, slot3)
	}
}
