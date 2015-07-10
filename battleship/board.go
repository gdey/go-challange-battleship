package battleship

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// EmptyBoard is the string template for a new board.
const EmptyBoard string = `   A B C D E F G H I J K L M N O P
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
  ╰─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─┴─╯`

// rowColForPosition will return a row and col for a position string. If an error occurs the row and col will be both set to zero.
func rowColForPosition(pos string) (row int, col int, err error) {
	if len(pos) > 3 || len(pos) < 2 {
		return 0, 0, errors.New("Bad position")
	}
	cstr := []rune(strings.ToUpper(pos[:1]))[0]
	rstr := pos[1:]
	col = int(cstr - 'A')
	row, err = strconv.Atoi(rstr)
	if err != nil {
		return 0, 0, err
	}
	row = row - 1
	if row > 16 || row < 0 {
		return 0, 0, errors.New("Bad Position")
	}
	return
}

// indexForPosition will return the corrosponding index for the position provided, if there is an error, err will not be nil, and the index will be zero.
func indexForPosition(pos string) (int, error) {
	r, c, err := rowColForPosition(pos)
	if err != nil {
		return 0, err
	}
	idx := indexForRowCol(r, c)
	return idx, nil
}

func indexForRowCol(r, c int) int {

	blength := strings.Index(EmptyBoard, "\n") + 2
	bRightPadding := strings.Index(EmptyBoard, "A")
	// Skip the first line (These are col labels) and move over to the right after row label
	startingOffset := (blength * 2) + bRightPadding
	rowOffset := (blength * 2 * r)
	colOffset := (c*2 - 1)
	return startingOffset + rowOffset + colOffset
}

func generateMap(hmarks []bmark, vmarks []bmark) string {
	newMap := []rune(EmptyBoard)
	blength := strings.Index(EmptyBoard, "\n") + 2

	for _, v := range hmarks {
		i := 0
		for _, rv := range v.name {
			if v.pos+1 >= len(newMap) {
				fmt.Print("index size is:", v.pos, len(newMap), "\n")
				continue
			}
			newMap[v.pos+i] = rv
			i++ // ranging over the string for runes does not mean that the i is going to match up
		}
	}
	for _, v := range vmarks {
		pos := v.pos
		for _, kv := range v.name {
			newMap[pos] = kv
			pos += blength
		}
	}
	return string(newMap)
}

type bmark struct {
	name string
	pos  int
}

// GameBoard represents the game's board. It should have enough information to store the state of the board so that it can be printed.
type GameBoard struct {
	vmarks []bmark
	hmarks []bmark
}

// Print writes the gameboard to the provided io.Writer.
func (b *GameBoard) Print(w io.Writer) error {
	s := generateMap(b.hmarks, b.vmarks)
	bytes := []byte(s)
	i, err := w.Write(bytes)
	if err != nil {
		return err
	}
	for len(bytes) > 0 && i < len(bytes) {
		bytes = bytes[:i]
		i, err = w.Write(bytes)
		if err != nil {
			return err
		}
	}
	return nil
}

// Place puts a mark onto the board. The mark is a string, and can be placed eighter vertically or horizontally.
func (b *GameBoard) Place(makr, position string, vertical bool) error {
	/* Fill in this Function */
	idx, err := indexForPosition(position)
	if err != nil {
		return err
	}
	if vertical {
		b.vmarks = append(b.vmarks, bmark{makr, idx})
		return nil
	}
	b.hmarks = append(b.hmarks, bmark{makr, idx})
	return nil
}

// Clear will remove any marks on the board and reset it to a clean state.
func (b *GameBoard) Clear() {
	b.vmarks = b.vmarks[0:0]
	b.hmarks = b.hmarks[0:0]
}

// NewGameBoard returns a new zero game board object, filled in with default values.
func NewGameBoard() *GameBoard {
	/* Fill in this Function */
	return &GameBoard{make([]bmark, 0), make([]bmark, 0)}
}
