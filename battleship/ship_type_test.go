package battleship

import "testing"

type ShipExpected struct {
	Name     string
	Expected string
}

func TestShipNames(t *testing.T) {
	TestCases := map[ShipType]ShipExpected{
		Missed:          ShipExpected{"Missed", "\u2573"},
		Hit:             ShipExpected{"Hit", "\u29BF"},
		AircraftCarrier: ShipExpected{"AircraftCarrier", "A*CARRIER"},
		Battleship:      ShipExpected{"Battleship", "B*SHIP*"},
		Submarine:       ShipExpected{"Submarine", "*SUB*"},
		Destroyer:       ShipExpected{"Destroyer", "DESTR"},
		Cruiser:         ShipExpected{"Cruiser", "CRUIS"},
		PatrolBoat:      ShipExpected{"PatrolBoat", "P*B"},
	}
	for ship, expected := range TestCases {
		if ship.Mark() != expected.Expected {
			t.Errorf("Expected %s to have string value of %s got %s instead.", expected.Name, expected.Expected, ship.Mark())
		}
	}

}
