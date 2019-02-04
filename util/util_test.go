package util

import (
	"testing"
	"github.com/ishiikurisu/wireworld"
)

func TestPackageCanLoadCsv(t *testing.T) {
	csv := `
		0;0;0;0;0
		0;1;3;3;0
		0;0;0;0;0
	`
	w, _ := LoadCsv(csv)
	if w.Dx != 5 {
		t.Error("Couldn't load X direction properly")
	}
	if w.Dy != 3 {
		t.Error("Couldn't load Y direction properly")
	}
	if w.World[0][0].Kind != wireworld.EMPTY {
		t.Error("Couldn't understand what an empty input is")	
	}
	if w.World[1][1].Kind != wireworld.POSITIVE {
		t.Error("Couldn't understand what an empty input is")	
	}
	if w.World[1][2].Kind != wireworld.CONDUCTOR {
		t.Error("Couldn't understand what an empty input is")	
	}
	if w.World[1][3].Kind != wireworld.CONDUCTOR {
		t.Error("Couldn't understand what an empty input is")	
	}
}