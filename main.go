package main

import (
	"bufio"
	"fmt"
	"os"

	"parking_lot_design/parking_lot"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please mention how would you like to give input.\nA. Enter file name\nB. Interactive command promt")
	isValid := false
	var usersChoice string
	var err error
	for !isValid {
		usersChoice, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading user's choice")
		} else {
			switch usersChoice[:1] {
			case "A":
				isValid = true
			case "B":
				isValid = true
			default:
				fmt.Println("Please enter either A or B depending upon your choice")
			}
		}
	}
	isCompleted := parking_lot.MainParkingLot(usersChoice[:1])
	if !isCompleted {
		fmt.Println("Unable to complete programs")
	}
}
