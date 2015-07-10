package battleship

import (
	"errors"
	"fmt"
	"io"
)

type gmark struct {
	name     string
	pos      string
	vertical bool
}

type gships struct {
	ship     ShipType
	pos      string
	vertical bool
	// Which positions have been hit so that we can calculate the hits, and do error checking
	hits []string
	// If the ship has been sank or not
	sank bool
	// This is what we will write to the board it's the precomuted set of marks to make for this ship.
	marks []gmark
}

func (g gships) PositionsForShip() (pos []string, err error) {
	pos, err = g.ship.PositionsForShip(g.pos, g.vertical)
	return
}

// G represents the game, and the state of the game. This is the main structure the run loop should interact with.
type G struct {
	Board        *GameBoard
	refreshBoard bool
	shipsSeen    map[ShipType]bool
	ships        []*gships
	misses       []string // These are the ones that were bad guesses
	seenPos      []string
	gameStarted  bool
	gameEnded    bool
	score        int
	round        int
	showBoard    bool
}

// ShowBoard toggles the printing to show all the ships.
func (g *G) ShowBoard() {
	g.showBoard = true
}

func (g *G) sunkenShips() (count int) {
	for _, s := range g.ships {
		if s.sank {
			count++
		}
	}
	return count
}

// NewGame returns a new zero game object, filled in with default values.
func NewGame() *G {
	gb := NewGameBoard()
	return &G{Board: gb}
}

func positionsIntersect(a []string, b []string) bool {
	for _, ap := range a {
		for _, bp := range b {
			if ap == bp {
				return true
			}
		}
	}
	return false
}

// PlaceShip allows the runner of the game to add ships to the game. The game must verify that it is not allowing ships to be placed that overlap.
func (g *G) PlaceShip(ship ShipType, position string, vertical bool) error {
	if g.gameStarted {
		return errors.New("Game already started can not add more ships")
	}
	_, err := indexForPosition(position)
	if err != nil {
		return err
	}
	_, ok := g.shipsSeen[ship]
	if ok {
		// We already have this type of ship don't allow it.
		return errors.New("This type of ship has already been add to the map")
	}
	if g.shipsSeen == nil {
		g.shipsSeen = make(map[ShipType]bool, 0)
	}
	g.shipsSeen[ship] = true
	ns := gships{
		pos:      position,
		vertical: vertical,
		ship:     ship,
	}
	if g.ships == nil {
		g.ships = make([]*gships, 0)
	}
	givenPos, err := ship.PositionsForShip(position, vertical)
	if err != nil {
		return err
	}
	// We need to now check to see if there are any interactions
	for _, s := range g.ships {
		gPos, err := s.PositionsForShip()
		if err != nil {
			// this should not happen.
			panic("We got an error in a trusted array!")
		}
		if positionsIntersect(givenPos, gPos) {
			return errors.New(fmt.Sprintf("Ships should not intersect: ship1: %v ship2: %v", givenPos, gPos))
		}
	}
	g.ships = append(g.ships, &ns)
	return nil
}

// Guess, allows the player to enter in the guesses for the round. The guess function must return an error of the incorrect number of guesses have been entered. It should, also, return which guesses hit something, and if any ships were sunk what the positions of those ships were.
func (g *G) Guess(positions []string) (hits []string, sank map[ShipType][]string, err error) {
	/* Fill in this Function */
	if len(positions) != 5 {
		return nil, nil, errors.New("There should be 5 guess each round.")
	}
	if g.seenPos == nil {
		g.seenPos = make([]string, 0)
	}
	if positionsIntersect(positions, g.seenPos) {
		return nil, nil, errors.New("Already seen one of the guesses.")
	}
	hits = make([]string, 0)
	misses := make([]string, 0)
	sank = make(map[ShipType][]string)
	// We need to go through the positions and see if any of them intersect
	for _, p := range positions {
		g.seenPos = append(g.seenPos, p)
		var didHit bool
		for _, gship := range g.ships {
			if gship.sank {
				continue // If the ship has been sank we don't need to look at it.
			}
			if gship.hits == nil {
				gship.hits = make([]string, 0)
			}
			if gship.marks == nil {
				gship.marks = make([]gmark, 0)
			}
			gshpos, _ := gship.PositionsForShip()
			for _, gshp := range gshpos {
				if p == gshp {
					g.refreshBoard = true
					didHit = true
					hits = append(hits, p)
					gship.hits = append(gship.hits, p)
					gship.sank = len(gship.hits) == gship.ship.Size()
					if gship.sank {
						gship.marks = []gmark{gmark{
							name:     gship.ship.Mark(),
							pos:      gship.pos,
							vertical: gship.vertical,
						}}
						sank[gship.ship] = gshpos
						g.score += gship.ship.Score()
					} else {
						gship.marks = append(gship.marks, gmark{
							name: Hit.Mark(),
							pos:  p,
						})
					}
				}
			}
		}
		if !didHit {
			misses = append(misses, p)
			g.refreshBoard = true
		}
	}
	if g.misses == nil {
		g.misses = make([]string, 0)
	}
	g.misses = append(g.misses, misses...)
	g.round += 1
	g.gameEnded = g.DidWin() || g.round >= 6
	return hits, sank, nil
}

// ShipsLeft returns which ships have not been sunk yet.
func (g *G) ShipsLeft() []ShipType {
	/* Fill in this Function */
	sl := make([]ShipType, 0)
	for _, s := range g.ships {
		if !s.sank {
			sl = append(sl, s.ship)
		}
	}
	return sl
}

// Score returns the current score for the game
func (g *G) Score() int {
	return g.score
}

// PrintBoard prints the representation of the baord to the given io.Writer.
func (g *G) PrintBoard(w io.Writer) error {
	if g.showBoard {
		g.refreshBoard = true
		g.showBoard = false
		g.Board.Clear()
		for _, s := range g.ships {
			g.Board.Place(s.ship.Mark(), s.pos, s.vertical)
		}
		return g.Board.Print(w)
	}
	if !g.refreshBoard {
		return g.Board.Print(w)
	}
	g.Board.Clear()
	// Need to place marks on the board.
	for _, m := range g.misses {
		g.Board.Place(Missed.Mark(), m, false)
	}
	for _, s := range g.ships {
		for _, m := range s.marks {
			g.Board.Place(m.name, m.pos, m.vertical)
		}
	}
	return g.Board.Print(w)
}

// Round return the current round the game is in.
func (g *G) Round() int {
	return g.round + 1
}

// TotalRounds returns the total number of rounds this game is allowed to go to.
func (g *G) TotalRounds() int {
	// We are going to hard code this for now.
	return 6
}

// Ended returns weather or not the game is still playable. Use the DidWin function to figure out if the User won or the Computer
func (g *G) Ended() bool {
	/* Fill in this Function */
	return g.gameEnded
}

// DidWin returns weather or not the user has won the game.
func (g *G) DidWin() bool {
	return len(g.ships) == g.sunkenShips()
}
