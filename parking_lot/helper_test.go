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

func TestGetRegNosForColor(t *testing.T) {
	parkingLotSize := 6
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
	parkingLot.ParkVehicle(vehicle1)
	vehicle2 := Vehicle{
		Color:     "Black",
		RegNumber: "KA-01-BB-0001",
	}
	parkingLot.ParkVehicle(vehicle2)
	vehicle3 := Vehicle{
		Color:     "White",
		RegNumber: "KA-01-HH-9999",
	}
	parkingLot.ParkVehicle(vehicle3)
	vehicle4 := Vehicle{
		Color:     "Red",
		RegNumber: "KA-01-HH-7777",
	}
	parkingLot.ParkVehicle(vehicle4)
	vehicle5 := Vehicle{
		Color:     "White",
		RegNumber: "KA-01-HH-8888",
	}
	parkingLot.ParkVehicle(vehicle5)
	vehicle6 := Vehicle{
		Color:     "White",
		RegNumber: "KA-01-DD-3333",
	}
	parkingLot.ParkVehicle(vehicle6)
	expectedRed := "KA-01-HH-7777"
	actualRed := parkingLot.GetRegNosForColor("Red")
	if actualRed != expectedRed {
		t.Errorf("Error GetRegNosForColor: Expected: %v Actual: %v", expectedRed, actualRed)
	}
	expectedWhite := "KA-01-HH-1234, KA-01-HH-9999, KA-01-HH-8888, KA-01-DD-3333"
	actualWhite := parkingLot.GetRegNosForColor("White")
	if expectedWhite != actualWhite {
		t.Errorf("Error GetRegNosForColor: Expected: %v Actual: %v", expectedWhite, actualWhite)
	}
	expectedBlue := NOT_FOUND
	actualBlue := parkingLot.GetRegNosForColor("Blue")
	if expectedBlue != actualBlue {
		t.Errorf("Error GetRegNosForColor: Expected: %v Actual: %v", expectedBlue, actualBlue)
	}
}

func TestGetSlotNosForColor(t *testing.T) {
	parkingLotSize := 4
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
	parkingLot.ParkVehicle(vehicle2)
	vehicle3 := Vehicle{
		Color:     "White",
		RegNumber: "KA-01-HH-9999",
	}
	slot3 := parkingLot.ParkVehicle(vehicle3)
	vehicle4 := Vehicle{
		Color:     "Red",
		RegNumber: "KA-01-HH-7777",
	}
	slot4 := parkingLot.ParkVehicle(vehicle4)
	expectedRed := util.IntToString(slot4)
	actualRed := parkingLot.GetSlotNosForColor("Red")
	if actualRed != expectedRed {
		t.Errorf("Error GetSlotNosForColor: Expected: %v Actual: %v", expectedRed, actualRed)
	}
	expectedWhite := util.IntToString(slot1) + ", " + util.IntToString(slot3)
	actualWhite := parkingLot.GetSlotNosForColor("White")
	if expectedWhite != actualWhite {
		t.Errorf("Error GetSlotNosForColor: Expected: %v Actual: %v", expectedWhite, actualWhite)
	}
}
