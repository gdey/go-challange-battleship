package battleship

type ShipType uint8

const (
	Missed ShipType = iota
	Hit
	AircraftCarrier
	Battleship
	Submarine
	Destroyer
	Cruiser
	PatrolBoat
)

func (st ShipType) Mark() string {
	s := "\u2573\u29BFA*CARRIERDESTR*SUB*SHIP*BCRUIS"
	switch st {
	case Missed:
		return s[0:3] //"\u2573"
	case Hit:
		return s[3:6] //"\u29BF"
	case AircraftCarrier:
		return s[6:15] //"A*CARRIER"
	case Destroyer:
		return s[15:20] //"DESTR"
	case Submarine:
		return s[20:25] //"*SUB*"
	case Battleship:
		return s[23:30] //"B*SHIP*"
	case PatrolBoat:
		return s[28:31] //"P*B"
	case Cruiser:
		return s[31:] //"CRUIS"
	default:
		return ""
	}
}
func (st ShipType) String() string {
	switch st {
	case Missed:
		return "missed"
	case Hit:
		return "hit"
	case AircraftCarrier:
		return "aircraft carrier"
	case Destroyer:
		return "destroyer"
	case Submarine:
		return "submarine"
	case Battleship:
		return "battleship"
	case PatrolBoat:
		return "patrol boat"
	case Cruiser:
		return "cruiser"
	default:
		return ""
	}
}
func (st ShipType) Size() int {
	switch st {
	case Missed:
		return 1
	case Hit:
		return 1
	case AircraftCarrier:
		return 5
	case Destroyer:
		return 3
	case Submarine:
		return 3
	case Battleship:
		return 4
	case PatrolBoat:
		return 2
	case Cruiser:
		return 3
	default:
		return 0
	}
}
func (st ShipType) Score() int {
	switch st {
	case AircraftCarrier:
		return 20
	case Destroyer:
		return 6
	case Submarine:
		return 6
	case Battleship:
		return 12
	case PatrolBoat:
		return 2
	case Cruiser:
		return 6
	default:
		return 0
	}
}
