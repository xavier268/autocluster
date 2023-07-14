package distance

// A distance matrix
// Optimised for storage efficiency. Zero value can be sused immediately.
type Matrix struct {
	data []float64 // stores the value as (0,1),     (0,2),(1,2),     (0,3),(1,3),(2,3),     (0,4),(1,4), 2,4),(3,4),    ...
}

// Get distance between i and j.
func (m *Matrix) Dist(i, j int) float64 {
	if i == j {
		return 0.
	}
	idx := index(i, j)
	if idx >= len(m.data) {
		return 0.
	}
	return m.data[idx]
}

func index(i, j int) int {
	if i < j {
		return i + j*(j-1)/2
	} else {
		return index(j, i)
	}
}

// Set a distance between i and j
func (m *Matrix) Set(i, j int, d float64) {
	if i == j {
		return
	}
	idx := index(i, j)
	if idx >= len(m.data) { // auto extend matrix when needed
		m.data = append(m.data, make([]float64, 1+idx-len(m.data))...)
	}
	m.data[index(i, j)] = d
}
