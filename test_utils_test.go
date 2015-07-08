package battleship_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const emptyBoard string = `   A B C D E F G H I J K L M N O P
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

func generateMap(hmarks []mark, vmarks []mark) string {
	newMap := []rune(emptyBoard)
	blength := strings.Index(emptyBoard, "\n") + 2

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

func StripFirstLastLine(str string) string {
	s := []rune(str)
	if s[0] == '\n' {
		s = s[1:]
	}
	if s[len(s)] == '\n' {
		s = s[0 : len(s)-1]
	}
	return string(s)
}

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

// If you know you positions are always valid, then you can use this function in single contex situtions.
func indexForPositionOrPanic(pos string) int {
	idx, err := indexForPosition(pos)
	if err != nil {
		panic(err)
	}
	return idx
}

func indexForRowCol(r, c int) int {

	blength := strings.Index(emptyBoard, "\n") + 2
	bRightPadding := strings.Index(emptyBoard, "A")
	// Skip the first line (These are col labels) and move over to the right after row label
	startingOffset := (blength * 2) + bRightPadding
	rowOffset := (blength * 2 * r)
	colOffset := (c*2 - 1)
	return startingOffset + rowOffset + colOffset
}

func fillPlacementsEverySpot(r rune) []Placement {
	str := string(r)
	placements := make([]Placement, 0, 256)
	for _, c := range colLabels {
		for _, r := range rowLabels {
			placements = append(placements, Placement{Mark: str, Position: c + r})
		}
	}
	return placements
}

func fillEverySpot(r rune) string {
	str := string([]rune{r})
	marks := make([]mark, 0, 256)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			marks = append(marks, mark{str, indexForRowCol(i, j)})
		}
	}
	return generateMap(marks, []mark{})
}
