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
