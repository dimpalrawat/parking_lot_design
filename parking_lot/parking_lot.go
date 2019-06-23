package parking_lot

func MainParkingLot(userChoice string) bool{
	if userChoice == "A" {
		return executeInputFromFile()
	} else if userChoice == "B" {
		return executeInputFromConsole()
	}
	return false
}
