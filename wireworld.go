package wireworld

type CellKind int
const (
	EMPTY CellKind = iota
	CONDUCTOR
	POSITIVE
	NEGATIVE
	CUSTOM
)

// The cell kind defines the basic unit. It implements a wireworld cell.
type Cell struct {
	// The cell kind determines how that cell should react
	Kind CellKind
	// In case of custom cells, a cell value can be used to determine special
	// reactions
	Value int
}

// this is the wireworld struct, where all the magic will happen
type Wireworld struct {
	// this is the length of the wireworld in the x direction
	Dx int
	// this is the length of the wireworld in the y direction
	Dy int
	// this is the matrix which wireworld will operate on
	World [][]Cell
}

// creates a new wireworld. all cells are initialized empty.
func NewWireworld(dx, dy int) Wireworld {
	w := Wireworld {
		Dx: dx,
		Dy: dy,
		World: make([][]Cell, dy),
	}

	for y := 0; y < dy; y++ {
		w.World[y] = make([]Cell, dx)
		for x := 0; x < dx; x++ {
			w.World[y][x] = Cell {
				Kind: EMPTY,
				Value: 0,
			}
		}
	}

	return w
}

// Generates a new wireworld by updating the given one
func Update(inlet Wireworld) Wireworld {
	outlet := NewWireworld(inlet.Dx, inlet.Dy)

	for y := 0; y < inlet.Dy; y++ {
		for x := 0; x < inlet.Dx; x++ {
			switch inlet.World[y][x].Kind {
			case EMPTY:
				outlet.World[y][x].Kind = EMPTY
				break
			case NEGATIVE:
				outlet.World[y][x].Kind = CONDUCTOR
				break
			case POSITIVE:
				outlet.World[y][x].Kind = NEGATIVE
				break
			case CONDUCTOR:
				fromX := max(0, x-1)
				toX := min(inlet.Dx, x+1)
				fromY := max(0, y-1)
				toY := min(inlet.Dy, y+1)
				heads := 0

				for i := fromX; i < toX; i++ {
					for j := fromY; j < toY; j++ {
						if inlet.World[j][i].Kind == POSITIVE {
							heads++
						}
					}
				}

				if heads == 1 || heads == 2 {
					outlet.World[y][x].Kind = POSITIVE
				} else {
					outlet.World[y][x].Kind = CONDUCTOR
				}
				break
			}
		}
	}

	return outlet
}
