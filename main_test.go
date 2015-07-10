package battleship_test

import (
	"bytes"
	"fmt"
	"testing"

	"./battleship"
)

// GuessTest is a guess in a game, and the expected values.
type GuessTest struct {
	// Guesses are the position guess that we are going to feed the game.
	Guesses []string
	// Hits are where we expects the hits to happen at in the order provided by the guesses
	Hits []string
	// Sank are the ships been sank and the the positions that lead to the sinking of the ship in the order they would occure on the board starting from the top right.
	Sank map[battleship.ShipType][]string
	// Map is the expect view of the map after the guess.
	Map string
	// Round is the expected round the game should be in after the guess.
	Round int
	// DidWin is weather the palery won after this guess
	DidWin bool
	// Err any expected error that this guess should have caused
	Err error
	// ShipsLeft are the ships that should be remaining in the game.
	ShipsLeft []battleship.ShipType
	// Score is the expected score after this guess
	Score int
}

// GameTest is the test for a full round of a game.
type GameTest struct {
	// The Description of the test. This makes it easier to identify a failing test.
	Description string
	// Placements are the starting positions of the Game Pieces
	Placements []ShipPlacement
	// Guesses are the guess that the player makes and the expect state of the game at that point.
	Guesses []GuessTest
}

// assertHits will check to make sure that the got hits array matches the expected array.
func assertHits(t *testing.T, label string, expected, got []string) {
	if len(expected) != len(got) {
		t.Errorf("%s array size does not match. %v (%v) != %v (%v) ", label, expected, len(expected), got, len(got))
		return // don't want to range over things if the sizes are not the same.
	}
	for i, val := range expected {
		if val != got[i] {
			t.Errorf("%s element %v does not match expected %v got %v", label, i, val, got[i])
		}
	}
}

// assertError checks to see that the error recieved is what was expected. It does not check to make sure the error acutally match, just that it's not nil if the expected is not nil.
func assertError(t *testing.T, label string, expected, got error) {
	if expected == nil {
		if got != nil {
			t.Errorf("%s was not expecting an error. Got: %v\n", label, got)
		}
		return
	}
	if got == nil {
		t.Errorf("%s was expecting an error. Did not get one.\n", label)
	}
	return
}
func testGame(label string, test GameTest, t *testing.T) {
	// Setup the game
	var buff bytes.Buffer
	game := battleship.NewGame()
	for _, placement := range test.Placements {
		err := game.PlaceShip(placement.Mark, placement.Position, placement.Vertical)
		assertError(t, label+" Placement ("+placement.Mark.String()+")", placement.Err, err)
	}
	if game.TotalRounds() != TotalRounds {
		t.Errorf("%sTotal rounds should be %v", label, TotalRounds)
	}
	err := game.PrintBoard(&buff)
	assertError(t, label+" PrintBoard", nil, err)
	if buff.String() != emptyBoard {
		if len(buff.String()) == 0 {
			t.Errorf("%sFailed initial print board test was expecting \n%s\nbut got nothing\n", label, emptyBoard)
		} else {
			t.Errorf("%sFailed initial print board test was expecting \n%s\ngot\n%s", label, emptyBoard, buff.String())
		}
	}
	buff.Reset()

	// Time to play the game.
	for _, guess := range test.Guesses {
		hits, sank, err := game.Guess(guess.Guesses)
		assertError(t, label+" Guess", guess.Err, err)
		assertHits(t, label+"Hits", guess.Hits, hits)
		for k, v := range guess.Sank {
			tb, ok := sank[k]
			if !ok {
				t.Errorf("%sSank: Expected to find ship %v", label, k)
				continue
			}
			assertHits(t, label+"Sank", v, tb)
		}
		d := game.Round()
		if d != guess.Round {
			t.Errorf("%sWas expecting round to be %v Got %v instead", label, guess.Round, d)
		}
		w := game.DidWin()
		if w != guess.DidWin {
			t.Errorf("%sWas expecting DidWin to be %v Got %v instead", label, guess.DidWin, w)
		}
		s := game.Score()
		if s != guess.Score {
			t.Errorf("%sWas expecting Score to be %v Got %v instead", label, guess.Score, s)
		}
		err = game.PrintBoard(&buff)
		if err != nil {
			t.Errorf("%sGot an unexpected error: %v", label, err)
		}
		if buff.String() != guess.Map {
			if len(buff.String()) == 0 {
				t.Errorf("%sFailed test was expecting \n%s\nbut got nothing\n", label, guess.Map)
			} else {
				t.Errorf("%sFailed test was expecting \n%s\ngot\n%s", label, guess.Map, buff.String())
			}
		}
		buff.Reset()
	}
}

func TestGames(t *testing.T) {
	var TestCases = []GameTest{
		GameTest{
			Description: "One ship board one hit",
			Placements: []ShipPlacement{
				ShipPlacement{
					Mark:     battleship.Submarine,
					Position: "C7",
					Vertical: true,
				},
			},
			Guesses: []GuessTest{
				GuessTest{
					Guesses: []string{"C6", "C8", "D7", "C7", "C9"},
					Hits:    []string{"C8", "C7", "C9"},
					Sank: map[battleship.ShipType][]string{
						battleship.Submarine: []string{"C7", "C8", "C9"},
					},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForRowCol(5, 2)},
							mark{"╳", indexForRowCol(6, 3)},
						},
						[]mark{
							mark{"*SUB*", indexForRowCol(6, 2)},
						},
					),
					Round:  2,
					DidWin: true,
					Score:  6,
				},
			},
		},
		GameTest{
			Description: "Two Ship board one hit",
			Placements: []ShipPlacement{
				ShipPlacement{
					Mark:     battleship.Submarine,
					Position: "C7",
					Vertical: true,
				},
				ShipPlacement{
					Mark:     battleship.Battleship,
					Position: "E5",
					Vertical: false,
				},
			},
			Guesses: []GuessTest{
				GuessTest{
					Guesses: []string{"C6", "C8", "D7", "C7", "C9"},
					Hits:    []string{"C8", "C7", "C9"},
					Sank: map[battleship.ShipType][]string{
						battleship.Submarine: []string{"C7", "C8", "C9"},
					},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForPositionOrPanic("C6")},
							mark{"╳", indexForPositionOrPanic("D7")},
						},
						[]mark{
							mark{"*SUB*", indexForPositionOrPanic("C7")},
						},
					),
					Round: 2,
					Score: 6,
				},
				GuessTest{
					Guesses: []string{"C1", "C2", "C3", "C4", "C5"},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForPositionOrPanic("C1")},
							mark{"╳", indexForPositionOrPanic("C2")},
							mark{"╳", indexForPositionOrPanic("C3")},
							mark{"╳", indexForPositionOrPanic("C4")},
							mark{"╳", indexForPositionOrPanic("C5")},
							mark{"╳", indexForPositionOrPanic("C6")},
							mark{"╳", indexForPositionOrPanic("D7")},
						},
						[]mark{
							mark{"*SUB*", indexForPositionOrPanic("C7")},
						},
					),
					Round: 3,
					Score: 6,
				},
				GuessTest{
					Guesses: []string{"E1", "E2", "E3", "E4", "E5"},
					Hits:    []string{"E5"},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForPositionOrPanic("C1")},
							mark{"╳", indexForPositionOrPanic("C2")},
							mark{"╳", indexForPositionOrPanic("C3")},
							mark{"╳", indexForPositionOrPanic("C4")},
							mark{"╳", indexForPositionOrPanic("C5")},
							mark{"╳", indexForPositionOrPanic("C6")},
							mark{"╳", indexForPositionOrPanic("D7")},
							mark{"╳", indexForPositionOrPanic("E1")},
							mark{"╳", indexForPositionOrPanic("E2")},
							mark{"╳", indexForPositionOrPanic("E3")},
							mark{"╳", indexForPositionOrPanic("E4")},
							mark{"⦿", indexForPositionOrPanic("E5")},
						},
						[]mark{
							mark{"*SUB*", indexForPositionOrPanic("C7")},
						},
					),
					Round: 4,
					Score: 6,
				},
				GuessTest{
					Guesses: []string{"D5", "F5", "E6", "M4", "N5"},
					Hits:    []string{"F5"},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForPositionOrPanic("C1")},
							mark{"╳", indexForPositionOrPanic("C2")},
							mark{"╳", indexForPositionOrPanic("C3")},
							mark{"╳", indexForPositionOrPanic("C4")},
							mark{"╳", indexForPositionOrPanic("C5")},
							mark{"╳", indexForPositionOrPanic("C6")},
							mark{"╳", indexForPositionOrPanic("D7")},
							mark{"╳", indexForPositionOrPanic("E1")},
							mark{"╳", indexForPositionOrPanic("E2")},
							mark{"╳", indexForPositionOrPanic("E3")},
							mark{"╳", indexForPositionOrPanic("E4")},
							mark{"\u29BF", indexForPositionOrPanic("E5")},
							mark{"\u29BF", indexForPositionOrPanic("F5")},
							mark{"╳", indexForPositionOrPanic("D5")},
							mark{"╳", indexForPositionOrPanic("E6")},
							mark{"╳", indexForPositionOrPanic("M4")},
							mark{"╳", indexForPositionOrPanic("N5")},
						},
						[]mark{
							mark{"*SUB*", indexForPositionOrPanic("C7")},
						},
					),
					Round: 5,
					Score: 6,
				},
				GuessTest{
					Guesses: []string{"P5", "O5", "H16", "M14", "N15"},
					Hits:    []string{},
					Map: generateMap(
						[]mark{
							mark{"╳", indexForPositionOrPanic("C1")},
							mark{"╳", indexForPositionOrPanic("C2")},
							mark{"╳", indexForPositionOrPanic("C3")},
							mark{"╳", indexForPositionOrPanic("C4")},
							mark{"╳", indexForPositionOrPanic("C5")},
							mark{"╳", indexForPositionOrPanic("C6")},
							mark{"╳", indexForPositionOrPanic("D7")},
							mark{"╳", indexForPositionOrPanic("E1")},
							mark{"╳", indexForPositionOrPanic("E2")},
							mark{"╳", indexForPositionOrPanic("E3")},
							mark{"╳", indexForPositionOrPanic("E4")},
							mark{"\u29BF", indexForPositionOrPanic("E5")},
							mark{"\u29BF", indexForPositionOrPanic("F5")},
							mark{"╳", indexForPositionOrPanic("D5")},
							mark{"╳", indexForPositionOrPanic("E6")},
							mark{"╳", indexForPositionOrPanic("M4")},
							mark{"╳", indexForPositionOrPanic("N5")},
							mark{"╳", indexForPositionOrPanic("P5")},
							mark{"╳", indexForPositionOrPanic("O5")},
							mark{"╳", indexForPositionOrPanic("H16")},
							mark{"╳", indexForPositionOrPanic("M14")},
							mark{"╳", indexForPositionOrPanic("N15")},
						},
						[]mark{
							mark{"*SUB*", indexForPositionOrPanic("C7")},
						},
					),
					Round: 6,
					Score: 6,
				},
			},
		},
	}
	for i, test := range TestCases {
		label := fmt.Sprintf("GameTest number %v: ", i)
		if len(test.Description) != 0 {
			label = fmt.Sprintf("For test %s: ", test.Description)
		}
		testGame(label, test, t)
	}
}
