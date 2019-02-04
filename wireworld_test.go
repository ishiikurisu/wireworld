package wireworld

import (
	"testing"
	"fmt"
)

func TestWireworldCanSetDefaultValues(t *testing.T) {
	w := NewWireworld(10, 10)
	if w.Dx != 10 {
		t.Error("WW couldn't set DX values")
	}
	if w.Dy != 10 {
		t.Error("WW couldn't set DY values")
	}
	for y := 0; y < w.Dy; y++ {
		for x := 0; x < w.Dx; x++ {
			if w.World[y][x].Kind != EMPTY {
				t.Error(fmt.Sprintf("Cell %d %d is not empty", x, y))
			}
		}
	}
}

func TestBasicWireworldOperation(t *testing.T) {
	csv := generateTestWorld()
	w, _ := LoadCsv(csv)
	w = Update(w)

	if w.World[0][0].Kind != EMPTY {
		t.Error("Couldn't understand what an empty cell is")
		return
	}
	if w.World[1][1].Kind != NEGATIVE {
		t.Error("Couldn't understand what a positive cell is")
		return
	}
	if w.World[1][2].Kind != POSITIVE {
		t.Error("Couldn't understand what a conductor cell is")
		return
	}
	if w.World[1][3].Kind != CONDUCTOR {
		t.Error("Couldn't understand what another conductor cell is")
		return
	}
}
