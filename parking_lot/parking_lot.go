package parking_lot

import (
	"fmt"
	"strings"
)

func MainParkingLot(commands []string) bool {
	parkingSlot := ParkingLot{
		RegToSlotNoMap: make(map[string]int),
		VacatedSlots:   &VacatedSlotsHeap{},
	}
	for commandNumber, command := range commands {
		commandWords := strings.Split(command, " ")
		if commandNumber == 0 {
			if CommandMap[Command(commandWords[0])] == 1 {
				slotSize := StringToInt(commandWords[1])
				parkingSlot.ParkingSlotSize = slotSize
				parkingSlot.VacatedSlots.InitializeHeap(slotSize)
				parkingSlot.BookedSlots = make([]*Vehicle, slotSize)
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
				slotNo := parkingSlot.ParkVehicle(vehicle)
				if slotNo == -1 {
					fmt.Println("Sorry, parking lot is full")
				} else {
					fmt.Println("Allocated slot number: " + IntToString(slotNo))
				}
			case 3:
				isVacated := parkingSlot.VacateParkingSpot(StringToInt(commandWords[1]))
				if isVacated {
					fmt.Println("Slot number " + commandWords[1] + " is free")
				} else {
					fmt.Println("No vehicle is parked on slot:" + commandWords[1])
				}
			case 4:
				//TODO: STATUS
			case 5:
				//TODO: REG NUMBERS FOR COLOR
			case 6:
				//TODO: SLOT NUMBER FOR REGISTRATION NUMBER
			case 7:
				//TODO: SLOT NUMBERS FOR COLOR

			default:
				fmt.Println("Please enter a valid command")
			}
		}
	}
	return true
}
