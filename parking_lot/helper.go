package parking_lot

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"

	"parking_lot_design/util"
	"parking_lot_design/util/file_handlers"
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

func executeInputFromFile() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please keep your input file inside parking_lot/static/input_files folder.")
	fmt.Print("Enter your file name: ")
	inputFileName, err := reader.ReadString('\n')
	if err != nil || inputFileName != "" && len(inputFileName) <2 {
		fmt.Println("Invalid input file name")
		return false
	} else {
		inputFileName = inputFileName[:len(inputFileName)-1]
	}
	file, err := file_handlers.GetFile(inputFileName)
	if err != nil || file == nil {
		fmt.Printf("Error in reading lines: %s", err)
		return false
	}
	defer file.Close()
	scanner := file_handlers.GetFileContentScanner(file)
	return executeParkingLot(scanner)
}

func executeInputFromConsole() bool {
	fmt.Println("Please enter your commands.")
	scanner := bufio.NewScanner(os.Stdin)
	return executeParkingLot(scanner)
}

func executeParkingLot(fileContentScanner *bufio.Scanner) bool {
	parkingSlot := ParkingLot{
		RegToSlotNoMap: make(map[string]int),
		VacatedSlots:   &VacatedSlotsHeap{},
	}
	commandNumber := 0
	for fileContentScanner.Scan() {
		commandNumber++
		command := fileContentScanner.Text()
		if command == "exit" {
			break
		}
		parkingSlot.executeCommand(command, commandNumber)
	}
	if commandNumber == 0 {
		return false
	}
	return true
}

func (this *ParkingLot) executeCommand (command string, commandNumber int) bool {
	commandWords := strings.Split(command, " ")
	if commandNumber == 1 {
		if CommandMap[Command(commandWords[0])] == 1 {
			slotSize := util.StringToInt(commandWords[1])
			this.ParkingSlotSize = slotSize
			this.VacatedSlots.InitializeHeap(slotSize)
			this.BookedSlots = make([]*Vehicle, slotSize)
			fmt.Println("Created a parking lot with " + commandWords[1] + " slots")
		} else {
			fmt.Println("Please initialize a parking lot first.")
			return false
		}
	} else {
		switch CommandMap[Command(commandWords[0])] {
		case 1:
			fmt.Println("This program currently support only one active parking slot at a time.")
			return false
		case 2:
			vehicle := Vehicle{
				Color:     commandWords[2],
				RegNumber: commandWords[1],
			}
			slotNo := this.ParkVehicle(vehicle)
			if slotNo == -1 {
				fmt.Println("Sorry, parking lot is full")
			} else {
				fmt.Println("Allocated slot number: " + util.IntToString(slotNo))
			}
		case 3:
			isVacated := this.VacateParkingSpot(util.StringToInt(commandWords[1]))
			if isVacated {
				fmt.Println("Slot number " + commandWords[1] + " is free")
			} else {
				fmt.Println("No vehicle is parked on slot: " + commandWords[1])
			}
		case 4:
			fmt.Println(this.PrintStatus())
		case 5:
			regNumbers := this.GetRegNosForColor(commandWords[1])
			fmt.Println(regNumbers)
		case 6:
			slotNumber := this.GetSlotNoFromRegNo(commandWords[1])
			fmt.Println(slotNumber)
		case 7:
			slotNumbers := this.GetSlotNosForColor(commandWords[1])
			fmt.Println(slotNumbers)
		default:
			fmt.Println("Please enter a valid command")
		}
	}
	return true
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

//Functions returns the current status of parking_lot
//Appending fixed number of spaces. Can add variable number of spaces also depending upon the slotNo
func (this *ParkingLot) PrintStatus() string {
	if this.VacatedSlots.Len() == this.ParkingSlotSize {
		return NO_VEHICLE_PARKED
	}
	resultString := "Slot No.     Registration No        Colour\n"
	for index, vehicle := range this.BookedSlots {
		if vehicle != nil {
			resultString = resultString + util.IntToString(index+1) + "            " + vehicle.RegNumber + "          " + vehicle.Color + "\n"
		}
	}
	resultLength := len(resultString)
	if resultLength < 2 {
		return resultString
	}
	return resultString[:resultLength-1]
}