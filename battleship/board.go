package battleship

import "io"

// GameBoard represents the game's board. It should have enough information to store the state of the board so that it can be printed.
type GameBoard struct{}

// Print writes the gameboard to the provided io.Writer.
func (b *GameBoard) Print(w io.Writer) error {
	/* Fill in this Function */
	return nil
}

// Place puts a mark onto the board. The mark is a string, and can be placed eighter vertically or horizontally.
func (b *GameBoard) Place(makr, position string, vertical bool) error {
	/* Fill in this Function */
	return nil
}

// NewGameBoard returns a new zero game board object, filled in with default values.
func NewGameBoard() *GameBoard {
	/* Fill in this Function */
	return &GameBoard{}
}
