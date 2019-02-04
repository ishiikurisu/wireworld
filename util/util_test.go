package util

import (
	"testing"
	"github.com/ishiikurisu/wireworld"
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
	if w.World[0][0].Kind != wireworld.EMPTY {
		t.Error("Couldn't understand what an empty cell is")
		return
	}
	if w.World[1][1].Kind != wireworld.POSITIVE {
		t.Error("Couldn't understand what a positive cell is")
		return
	}
	if w.World[1][2].Kind != wireworld.CONDUCTOR {
		t.Error("Couldn't understand what a conductor cell is")
		return
	}
	if w.World[1][3].Kind != wireworld.CONDUCTOR {
		t.Error("Couldn't understand what another conductor cell is")
		return
	}
}
