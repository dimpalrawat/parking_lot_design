package parking_lot

type Command string

type Vehicle struct {
	Color     string
	RegNumber string
}

type ParkingLot struct {
	ParkingSlotSize  int
	VacatedSlots     *VacatedSlotsHeap
	BookedSlots      []*Vehicle
	RegToSlotNoMap   map[string]int
}

type Slot struct {
	SlotNo  int
	Vehicle *Vehicle
}

type ParkingLotInterface interface {
	ParkVehicle(vehicle Vehicle) int
	VacateParkingSpot(spotNumber int) bool
	GetSlotNoFromRegNo(registrationNo string) string
	GetSlotNosForColor(color string) string
	GetRegNosForColor(color string) string
	PrintStatus() string
}