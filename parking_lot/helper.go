package parking_lot

import (
	"container/heap"
	"parking_lot_design/util"
)

type VacatedSlotsHeap []int

func (h VacatedSlotsHeap) Len() int           { return len(h) }
func (h VacatedSlotsHeap) Less(i, j int) bool { return (h[i]) < h[j] }
func (h VacatedSlotsHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *VacatedSlotsHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *VacatedSlotsHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Initializing MinHeap
func (h *VacatedSlotsHeap) InitializeHeap(n int) {
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
}

//Function takes the vehicle object and returns slotNo if available else -1
func (this *ParkingLot) ParkVehicle(vehicle Vehicle) int {
	if this.VacatedSlots.Len() == 0 {
		return -1
	}
	slotNo := heap.Pop(this.VacatedSlots)
	intSlotNo := slotNo.(int)
	this.RegToSlotNoMap[vehicle.RegNumber] = intSlotNo
	this.BookedSlots[intSlotNo-1] = &vehicle
	return slotNo.(int)
}

//Function vacates the parking slot and returns true if vehicle is present else returns false
func (this *ParkingLot) VacateParkingSpot(spotNumber int) bool {
	if this.BookedSlots[spotNumber-1] == nil {
		return false
	}
	heap.Push(this.VacatedSlots, spotNumber)
	vechicle := this.BookedSlots[spotNumber-1]
	delete(this.RegToSlotNoMap, vechicle.RegNumber)
	this.BookedSlots[spotNumber-1] = nil
	return true
}

//Function returns slotNo for given registrationNo
func (this *ParkingLot) GetSlotNoFromRegNo(registrationNo string) string {
	slotNo, ok := this.RegToSlotNoMap[registrationNo]
	if !ok {
		return NOT_FOUND
	}
	return util.IntToString(slotNo)
}

//Function returns list of registrationNos for given color
func (this *ParkingLot) GetRegNosForColor(color string) string {
	resultString := ""
	for _, vehicle := range this.BookedSlots {
		if vehicle != nil && vehicle.Color == color {
			resultString = resultString + vehicle.RegNumber + ", "
		}
	}
	resultStrLen := len(resultString)
	if resultStrLen < 2 {
		return NOT_FOUND
	}
	return resultString[:resultStrLen-2]
}

//Function returns list of slotNos for given color
func (this *ParkingLot) GetSlotNosForColor(color string) string {
	resultString := ""
	for index, vehicle := range this.BookedSlots {
		if vehicle != nil && vehicle.Color == color {
			resultString = resultString + util.IntToString(index+1) + ", "
		}
	}
	resultStrLen := len(resultString)
	if resultStrLen < 2 {
		return NOT_FOUND
	}
	return resultString[:resultStrLen-2]
}