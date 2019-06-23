package parking_lot

import (
	"bufio"
	"strings"
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

func TestPrintStatus(t *testing.T) {
	parkingLotSize := 4
	parkingLot := ParkingLot{
		RegToSlotNoMap:   make(map[string]int),
		BookedSlots:      make([]*Vehicle, parkingLotSize),
		VacatedSlots:     &VacatedSlotsHeap{},
	}
	parkingLot.ParkingSlotSize = parkingLotSize
	parkingLot.VacatedSlots.InitializeHeap(parkingLotSize)
	expectedStatus := "Novehicleisparked"
	actualStatus := parkingLot.PrintStatus()
	actualStatus = util.StripSpaces(util.StripNewLines(actualStatus))
	if actualStatus != expectedStatus {
		t.Errorf("Error GetSlotNosForColor: Expected: %v Actual: %v", expectedStatus, actualStatus)
	}
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
	slot3 := parkingLot.ParkVehicle(vehicle3)
	vehicle4 := Vehicle{
		Color:     "Red",
		RegNumber: "KA-01-HH-7777",
	}
	parkingLot.ParkVehicle(vehicle4)
	expectedStatus = "SlotNo.RegistrationNoColour1KA-01-HH-1234White2KA-01-BB-0001Black3KA-01-HH-9999White4KA-01-HH-7777Red"
	actualStatus = parkingLot.PrintStatus()
	actualStatus = util.StripSpaces(util.StripNewLines(actualStatus))
	if actualStatus != expectedStatus {
		t.Errorf("Error GetSlotNosForColor: Expected: %v Actual: %v", expectedStatus, actualStatus)
	}
	parkingLot.VacateParkingSpot(slot3)
	expectedStatus = "SlotNo.RegistrationNoColour1KA-01-HH-1234White2KA-01-BB-0001Black4KA-01-HH-7777Red"
	actualStatus = parkingLot.PrintStatus()
	actualStatus = util.StripSpaces(util.StripNewLines(actualStatus))
	if expectedStatus != actualStatus {
		t.Errorf("Error GetSlotNosForColor: Expected: %v Actual: %v", expectedStatus, actualStatus)
	}
}

func TestExecuteParkingLot(t *testing.T) {
	strInput1 := `park KA-01-HH-1233 White`
	scanner := bufio.NewScanner(strings.NewReader(strInput1))
	strReader := strings.NewReader(strInput1)
	strReader.ReadByte()
	result1 := executeParkingLot(scanner)
	if result1 != false {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", false, result1)
	}
	strInput2 := ``
	scanner = bufio.NewScanner(strings.NewReader(strInput2))
	result2 := executeParkingLot(scanner)
	if result2 != false {
		t.Errorf("Error parkingLot 2: Expected: %v Actual: %v ", false, result2)
	}
	strInput3 := `create_parking_lot 6
park KA-01-HH-1233 White
park KA-01-HH-9999 White
park KA-01-BB-0001 Black
park KA-01-HH-7777 Red
park KA-01-HH-2701 Blue
park KA-01-HH-3141 Black
leave 4
status
park KA-01-P-333 White
park DL-12-AA-9999 White
registration_numbers_for_cars_with_colour White
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-HH-3141
slot_number_for_registration_number MH-04-AY-1111`
	scanner = bufio.NewScanner(strings.NewReader(strInput3))
	result3 := executeParkingLot(scanner)
	if result3 != true {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", true, result3)
	}
	strInput4 := `create_parking_lot 4
park KA-01-HH-1233 White
park KA-01-HH-9999 Black
create_parking_lot 4
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-HH-1233
slot_number_for_registration_number MH-04-AY-1111`
	scanner = bufio.NewScanner(strings.NewReader(strInput4))
	result4 := executeParkingLot(scanner)
	if result4 != true {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", true, result4)
	}
	strInput5 := `create_parking_lot 4
leave 3
`
	scanner = bufio.NewScanner(strings.NewReader(strInput5))
	result5:= executeParkingLot(scanner)
	if result5 != true {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", true, result5)
	}
	strInput6 := `


`
	scanner = bufio.NewScanner(strings.NewReader(strInput6))
	result6:= executeParkingLot(scanner)
	if result6 != false {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", false, result6)
	}
	strInput7 := `
park KA-01-HH-1233 White
park KA-01-HH-9999 Black
leave 5
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-HH-1233
`
	scanner = bufio.NewScanner(strings.NewReader(strInput7))
	result7:= executeParkingLot(scanner)
	if result6 != false {
		t.Errorf("Error parkingLot: Expected: %v Actual: %v ", false, result7)
	}
}