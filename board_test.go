package battleship_test

import (
	"bytes"
	"testing"

	"./battleship"
)

const (
	TotalRounds = 6
)

// Placement define where to place a label, these generally represent a ship.
type Placement struct {
	// The string to print onto the map.
	Mark string
	// Position on the map, should be in the format of [A-P][1-16]
	Position string
	Vertical bool
}

type ShipPlacement struct {
	// The type of ship being placed onto the map
	Mark battleship.ShipType
	// Position is where on the map it is located, it should be in the format of [A-P][1-16]
	Position string
	// Is the mark vertical or horizontally laid out.
	Vertical bool
	// The expected error if any for this placement.
	Err error
}

// BoardTest defines a test for the baord.
type BoardTest struct {
	// Map is how the board should look like; this is what the generated map is compared to.
	Map string
	// Placements are the where the ship and other labels are located. This is what the test is given.
	Placements []Placement
}

var colLabels []string
var rowLabels []string

func init() {
	colLabels = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
	rowLabels = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}
}

type mark struct {
	name string
	pos  int
}

func BoardPrintTest(test BoardTest, t *testing.T) {
	var buff bytes.Buffer
	b := battleship.NewGameBoard()
	for _, placement := range test.Placements {
		b.Place(placement.Mark, placement.Position, placement.Vertical)
	}
	err := b.Print(&buff)
	if err != nil {
		t.Error("Got an unexpected error.", err)
	}
	if buff.String() != test.Map {
		if len(buff.String()) == 0 {
			t.Errorf("Failed test was expecting \n%s\nbut got nothing\n", test.Map)
		} else {
			t.Errorf("Failed test was expecting \n%s\ngot\n%s", test.Map, buff.String())
		}
	}
	buff.Reset()
}

func TestBattleshipBoard(t *testing.T) {
	var TestCases = []BoardTest{
		BoardTest{Map: emptyBoard},
		BoardTest{
			Map: generateMap([]mark{
				mark{"\u29BF", indexForPositionOrPanic("A1")},
			}, []mark{}),
			Placements: []Placement{
				Placement{
					Mark:     "\u29BF",
					Position: "A1",
				},
			},
		},
		BoardTest{
			Map: generateMap([]mark{
				mark{"\u29BF", indexForPositionOrPanic("A2")},
			}, []mark{}),
			Placements: []Placement{
				Placement{
					Mark:     "\u29BF",
					Position: "A2",
				},
			},
		},
		BoardTest{
			Map: generateMap([]mark{
				mark{"\u29BF", indexForPositionOrPanic("I8")},
			}, []mark{}),
			Placements: []Placement{
				Placement{
					Mark:     "\u29BF",
					Position: "I8",
				},
			},
		},
		BoardTest{
			Map: generateMap([]mark{
				mark{"\u29BF", indexForPositionOrPanic("P16")},
			}, []mark{}),
			Placements: []Placement{
				Placement{
					Mark:     "\u29BF",
					Position: "P16",
				},
			},
		},
		BoardTest{
			Map:        fillEverySpot('\u29BF'),
			Placements: fillPlacementsEverySpot('\u29BF'),
		},
		BoardTest{
			Map: generateMap([]mark{
				mark{"*SUB*", indexForPositionOrPanic("A1")},
			}, []mark{}),
			Placements: []Placement{
				Placement{
					Mark:     "*SUB*",
					Position: "A1",
				},
			},
		},
		BoardTest{
			Map: generateMap([]mark{}, []mark{
				mark{"*SUB*", indexForPositionOrPanic("A1")},
			}),
			Placements: []Placement{
				Placement{
					Mark:     "*SUB*",
					Position: "A1",
					Vertical: true,
				},
			},
		},
		BoardTest{
			Map: generateMap([]mark{}, []mark{
				mark{"*SUB*", indexForPositionOrPanic("B7")},
			}),
			Placements: []Placement{
				Placement{
					Mark:     "*SUB*",
					Position: "B7",
					Vertical: true,
				},
			},
		},
	}
	for _, test := range TestCases {
		BoardPrintTest(test, t)
	}
	return
}
