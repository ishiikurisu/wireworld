package util

import (
	"github.com/ishiikurisu/wireworld"
	"strings"
)

// loads
func LoadCsv(csv string) (wireworld.Wireworld, error) {
	// building string matrix
	csv = strings.TrimSpace(csv)
	raw := make([][]string, 0)
	lines := strings.Split(csv, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fields := strings.Split(line, ";")
		raw = append(raw, fields)
	}

	// storing data in wireworld object
	dx := len(raw[0])
	dy := len(raw)
	w := wireworld.NewWireworld(dx, dy)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			switch raw[y][x] {
			case "1":
				w.World[y][x].Kind = wireworld.POSITIVE
				break
			case "3":
				w.World[y][x].Kind = wireworld.CONDUCTOR
				break
			}
		}
	}


	return w, nil
}
