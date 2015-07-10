package battleship

var rowlabel, collabel []string

func init() {
	rowlabel = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16"}
	collabel = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
}

func (st ShipType) PositionsForShip(p string, vertical bool) ([]string, error) {
	r, c, err := rowColForPosition(p)
	if err != nil {
		return nil, err
	}
	pos := []string{p}
	var length int
	if vertical {
		length = len(rowlabel)
	} else {
		length = len(collabel)
	}
	for i := 1; i < st.Size(); i++ {
		var np int
		if vertical {
			np = (r + i) % length
			pos = append(pos, collabel[c]+rowlabel[np])
		} else {
			np = (c + i) % length
			pos = append(pos, collabel[np]+rowlabel[r])
		}
	}
	return pos, nil
}
