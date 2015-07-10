package battleship

import "io"

// NewGame returns a new zero game object, filled in with default values.
func NewGame() *G {
	/* Fill in this Function */
	return &G{}
}

// G represents the game, and the state of the game. This is the main structure the run loop should interact with.
type G struct{}

// PlaceShip allows the runner of the game to add ships to the game. The game must verify that it is not allowing ships to be placed that overlap.
func (g *G) PlaceShip(ship ShipType, position string, vertical bool) error {
	/* Fill in this Function */
	return nil
}

// Guess, allows the player to enter in the guesses for the round. The guess function must return an error of the incorrect number of guesses have been entered. It should, also, return which guesses hit something, and if any ships were sunk what the positions of those ships were.
func (g *G) Guess(positions []string) (hits []string, sank map[ShipType][]string, err error) {
	/* Fill in this Function */
	return
}

// ShipsLeft returns which ships have not been sunk yet.
func (g *G) ShipsLeft() []ShipType {
	/* Fill in this Function */
	return []ShipType{}
}

// Score returns the current score for the game
func (g *G) Score() int {
	/* Fill in this Function */
	return 0
}

// PrintBoard prints the representation of the baord to the given io.Writer.
func (g *G) PrintBoard(w io.Writer) error {
	/* Fill in this Function */
	return nil
}

// Round return the current round the game is in.
func (g *G) Round() int {
	/* Fill in this Function */
	return 1
}

// TotalRounds returns the total number of rounds this game is allowed to go to.
func (g *G) TotalRounds() int {
	/* Fill in this Function */
	return 0
}

// Ended returns weather or not the game is still playable. Use the DidWin function to figure out if the User won or the Computer
func (g *G) Ended() bool {
	/* Fill in this Function */
	return false
}

// DidWin returns weather or not the user has won the game.
func (g *G) DidWin() bool {
	/* Fill in this Function */
	return false
}

// ShowBoard tells the game object that we want to show the board with all the ships revealed.
func (g *G) ShowBoard() {
	/* Fill in this Function */
	return
}
