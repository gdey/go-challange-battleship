package battleship

import "io"

/* Board represents the game board. It allows one to print the board out to any io.Writer.
A board should be printed as the following:

# A New board should look like this:
   A B C D E F G H I J K L M N O P
  ╭─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─╮
 1│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 2│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 3│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 4│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 5│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 6│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 7│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 8│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 9│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
10│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
11│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
12│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
13│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
14│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
15│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
16│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ╰─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─╯

# Here a player has selected a few hits, and a miss.
   A B C D E F G H I J K L M N O P
  ╭─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─╮
 1│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 2│ │ │ │ │ │ │⦿│⦿│ │ │⦿│ │⦿│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 3│ │ │ │ │ │ │ │ │ │ │ │ │⦿│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 4│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 5│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 6│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 7│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 8│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 9│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
10│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
11│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
12│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
13│ │ │ │ │ │ │ │╳│ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
14│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
15│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
16│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ╰─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─╯

  # A Board with hits on every ships on E2:H2,K2:K4,M2:M5,E5:I5,D8:D10, and H9:H10,  as well as a miss on H13  would look like this:
   A B C D E F G H I J K L M N O P
  ╭─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─╮
 1│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 2│ │ │ │ │ │*SUB*│ │ │D│ │B│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼E┼─┼*┼─┼─┼─┤
 3│ │ │ │ │ │ │ │ │ │ │S│ │S│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼T┼─┼H┼─┼─┼─┤
 4│ │ │ │ │ │ │ │ │ │ │R│ │I│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼P┼─┼─┼─┤
 5│ │ │ │ │A*CARRIER│ │ │ │*│ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 6│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 7│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 8│ │ │ │C│ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼R┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
 9│ │ │ │U│ │ │ │P│ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼I┼─┼─┼─┼*┼─┼─┼─┼─┼─┼─┼─┼─┤
10│ │ │ │S│ │ │ │B│ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
11│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
12│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
13│ │ │ │ │ │ │ │╳│ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
14│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
15│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┼─┤
16│ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │ │
  ╰─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─╯

*/

const (
	LeftTopCorner                rune = '\u256D'
	RightBottomCorner            rune = '\u2570'
	VerticalBar                  rune = '\u2502'
	HorizontalBar                rune = '\u2500'
	RightTopCorner               rune = '\u256E'
	LeftBottomCorner             rune = '\u256F'
	HorizontalMidUpVerticalBar   rune = '\u2534'
	CrossBar                     rune = '\u253C'
	VirticalMidLeftHorizontalBar rune = '\u2524'
	XMark                        rune = '\u2573'
	OMark                        rune = '\u29BF'
)

// Board interface is the supported behavior for a baord.
type Board interface {
	// Print allows the board to be printed to the given io.Writer object.
	Print(w io.Writer) error
	// Place allow the baord to have different strings placed on the board.
	Place(mark, position string, vertical bool) error
}

// Game interface is the supported behavior of a game of battleship.
type Game interface {
	// PlaceShip places the ship into the game a the location. The game will make sure that the ship being place does not cross any other ships placed in the game.
	PlaceShip(ship ShipType, position string, vertical bool) error
	// Guess allows the user to enter in a series of Guesses, and returns where the hits were, and what ships (if any) were sank.
	Guess(positions []string) (hits []string, sank map[ShipType][]string, err error)
	// ShipsLeft returns the ships that are still alive
	ShipsLeft() []ShipType
	// Score returns the users current score.
	Score() int
	// PrintBoard prints the board to the provided io.Writer.
	PrintBoard(w io.Writer) error
	// Round returns which round the game is currently at.
	Round() int
	// TotalRounds returns the total number of rounds in the game.
	TotalRounds() int
	// DidWin returns weather or not the user won the game
	DidWin() bool
	// Ended returns weather or not the games has ended. If the game has ended, entering more guesses will resulted in a error.
	Ended() bool
	// ShowBoard tells the game object that we want to show the board with all the ships revealed.
	ShowBoard()
}
