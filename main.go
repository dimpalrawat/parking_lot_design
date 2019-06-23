package main

import "parking_lot_design/parking_lot"

func main() {
	commands := []string{"create_parking_lot 6","park KA-01-HH-1234 White","park KA-01-BB-0001 Black","park KA-01-HH-7777 Red","leave 2","park KA-01-HH-2701 Blue"}
	parking_lot.MainParkingLot(commands)
}
