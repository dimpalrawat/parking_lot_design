package parking_lot

type Command string

type Vehicle struct {
	Color     string
	RegNumber string
}

type ParkingLot struct {
	ParkingSlotSize int
	VacatedSlots    *VacatedSlotsHeap
	Slots           []*Vehicle
	RegToSlotNoMap  map[string]int
}

//Could have used a separate struct for slot
/*type Slot struct {
	SlotNo  int
	Vehicle *Vehicle
}*/

type ParkingLotInterface interface {
	ParkVehicle(vehicle Vehicle) int
	VacateParkingSpot(spotNumber int) bool
	GetSlotNoFromRegNo(registrationNo string) string
	GetSlotNosForColor(color string) string
	GetRegNosForColor(color string) string
	PrintStatus() string
}