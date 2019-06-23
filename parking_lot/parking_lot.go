package parking_lot

func MainParkingLot(userChoice string) bool{
	if userChoice == "A" {
		return executeInputFromFile()
	} else {
		return executeInputFromConsole()
	}
	return false
}
