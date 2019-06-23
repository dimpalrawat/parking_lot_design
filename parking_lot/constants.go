package parking_lot

const (
	CREATE_PARKING_LOT Command = "create_parking_lot"
	PARK               Command = "park"
	LEAVE              Command = "leave"
	STATUS             Command = "status"
	RES_NOS_FOR_COLOR  Command = "registration_numbers_for_cars_with_colour"
	SLOT_NO_FOR_REG_NO Command = "slot_number_for_registration_number"
	SLOT_NOS_FOR_COLOR Command = "slot_numbers_for_cars_with_colour"
)

var CommandMap = map[Command]int{
	CREATE_PARKING_LOT: 1,
	PARK:               2,
	LEAVE:              3,
	STATUS:             4,
	RES_NOS_FOR_COLOR:  5,
	SLOT_NO_FOR_REG_NO: 6,
	SLOT_NOS_FOR_COLOR: 7,
}

const (
	NOT_FOUND = "Not found"
	NO_VEHICLE_PARKED = "No vehicle is parked"
)