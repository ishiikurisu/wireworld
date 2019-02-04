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
