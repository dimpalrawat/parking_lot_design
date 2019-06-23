package main

import (
	"fmt"

	"parking_lot_design/parking_lot"
)

func main() {
	inputFileName := "commands.txt"
	isCompleted := parking_lot.MainParkingLot(inputFileName)
	if !isCompleted {
		fmt.Println("Unable to complete programs")
	}
}
