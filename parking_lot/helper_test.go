package parking_lot

import (
	"testing"
	
	"parking_lot_design/util"
)

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

func TestVacateParkingSpot(t *testing.T) {
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
	vehicle2 := Vehicle{
		Color:     "Black",
		RegNumber: "KA-01-BB-0001",
	}
	slot2 := parkingLot.ParkVehicle(vehicle2)
	isVacated := parkingLot.VacateParkingSpot(slot1)
	if isVacated != true {
		t.Errorf("Error VacateParkingSpot: Expected: %v Actual: %v", true, isVacated)
	}
	isVacated = parkingLot.VacateParkingSpot(slot2)
	if isVacated != true {
		t.Errorf("Error VacateParkingSpot: Expected: %v Actual: %v", true, isVacated)
	}
	isVacated = parkingLot.VacateParkingSpot(slot2)
	if isVacated != false {
		t.Errorf("Error VacateParkingSpot: Expected: %v Actual: %v", false, isVacated)
	}
}

func TestGetSlotNoFromRegNo(t *testing.T) {
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
	vehicle2 := Vehicle{
		Color:     "Black",
		RegNumber: "KA-01-BB-0001",
	}
	slot2 := parkingLot.ParkVehicle(vehicle2)
	slotNumber1 := util.StringToInt(parkingLot.GetSlotNoFromRegNo(vehicle1.RegNumber))
	if slotNumber1 == 0 || slotNumber1 != slot1 {
		t.Errorf("Error GetSlotNoFromRegNo: Expected: %v Actual: %v", slot1, slotNumber1)
	}
	slotNumber2 := util.StringToInt(parkingLot.GetSlotNoFromRegNo(vehicle2.RegNumber))
	if slotNumber2 == 0 || slotNumber2 != slot2 {
		t.Errorf("Error GetSlotNoFromRegNo: Expected: %v Actual: %v", slot2, slotNumber2)
	}
	slotNumber3 := util.StringToInt(parkingLot.GetSlotNoFromRegNo("KA-01-CC-3241"))
	if slotNumber3 != 0 {
		t.Errorf("Error GetSlotNoFromRegNo: Expected: %v Actual: %v", 0, slotNumber3)
	}
}