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

func (st ShipType) String() string {
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
