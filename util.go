package wireworld

import (
	"strings"
)

// loads
func LoadCsv(csv string) (Wireworld, error) {
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
	w := NewWireworld(dx, dy)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			switch raw[y][x] {
			case "1":
				w.World[y][x].Kind = POSITIVE
				break
			case "2":
				w.World[y][x].Kind = NEGATIVE
				break
			case "3":
				w.World[y][x].Kind = CONDUCTOR
				break
			}
		}
	}


	return w, nil
}

func ConvertToCsv(w Wireworld) string {
	s := ""

	for y := 0; y < w.Dy; y++ {
		for x := 0; x < w.Dx; x++ {
			switch w.World[y][x].Kind {
			case EMPTY:
				s += "0"
				break
			case NEGATIVE:
				s += "2"
				break
			case POSITIVE:
				s += "1"
				break
			case CONDUCTOR:
				s += "3"
				break
			}
			if x == w.Dx-1 {
				s += "\n"
			} else {
				s += ";"
			}
		}
	}

	return s
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
