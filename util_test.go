package wireworld

import (
	"testing"
	"fmt"
)

func generateTestWorld() string {
	return `
		0;0;0;0;0
		0;1;3;3;0
		0;0;0;0;0
	`
}

func TestPackageCanLoadCsv(t *testing.T) {
	csv := generateTestWorld()
	w, _ := LoadCsv(csv)
	if w.Dx != 5 {
		t.Error("Couldn't load X direction properly")
		return
	}
	if w.Dy != 3 {
		t.Error("Couldn't load Y direction properly")
		return
	}
	if w.World[0][0].Kind != EMPTY {
		t.Error("Couldn't understand what an empty cell is")
		return
	}
	if w.World[1][1].Kind != POSITIVE {
		t.Error("Couldn't understand what a positive cell is")
		return
	}
	if w.World[1][2].Kind != CONDUCTOR {
		t.Error("Couldn't understand what a conductor cell is")
		return
	}
	if w.World[1][3].Kind != CONDUCTOR {
		t.Error("Couldn't understand what another conductor cell is")
		return
	}
}

func TestPackageCanConvertToCsv(t *testing.T) {
	csv := generateTestWorld()
	w, _ := LoadCsv(csv)
	w = Update(w)
	w = Update(w)
	maybe := ConvertToCsv(w)
	target := "0;0;0;0;0\n" +
	          "0;3;2;1;0\n" +
	          "0;0;0;0;0\n"
	if maybe != target {
		fmt.Println("expected:\n")
		fmt.Println(target)
		fmt.Println("reality:\n")
		fmt.Println(maybe)
		t.Error("Couldn't convert world to CSV")
		return
	}
}
